package tokengenerator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type hashGenerator struct {
	key []byte
}

func (p *hashGenerator) generateHash(txt string) (hashData string, err error) {
	block, err := aes.NewCipher(p.key)
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

	ciphertext := gcm.Seal(nonce, nonce, []byte(txt), nil)
	return hex.EncodeToString(ciphertext), nil
}

func (p *hashGenerator) parseHash(hashData string) (txt string, err error) {
	ciphertext, err := hex.DecodeString(hashData)
	if err != nil {
		return txt, err
	}

	block, err := aes.NewCipher(p.key)
	if err != nil {
		return txt, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return txt, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return txt, err
	}

	return string(plaintext), nil
}

func newHashGenerator(key []byte) *hashGenerator {
	return &hashGenerator{key: key}
}
