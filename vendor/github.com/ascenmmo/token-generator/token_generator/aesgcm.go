package tokengenerator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"io"
	"time"
)

type aesgmGenerator struct {
	key []byte
}

type aesgmClaims struct {
	TTL  int64 `json:"dead_at"`
	Info tokentype.Info
}

func (a *aesgmGenerator) generate(info tokentype.Info) (token string, err error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	claims := aesgmClaims{
		Info: info,
		TTL:  time.Now().Add(info.TTL).Unix(),
	}

	marshal, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, marshal, nil)
	return hex.EncodeToString(ciphertext), nil
}

func (a *aesgmGenerator) parse(token string) (info tokentype.Info, err error) {
	var claims aesgmClaims
	ciphertext, err := hex.DecodeString(token)
	if err != nil {
		return info, err
	}

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return info, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return info, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return info, err
	}

	err = json.Unmarshal(plaintext, &claims)
	if err != nil {
		return info, err
	}

	if claims.TTL < time.Now().Unix() {
		return info, errors.New("token expired")
	}
	info = claims.Info
	return info, nil
}

func newAESGMGenerator(key []byte) *aesgmGenerator {
	return &aesgmGenerator{key: key}
}
