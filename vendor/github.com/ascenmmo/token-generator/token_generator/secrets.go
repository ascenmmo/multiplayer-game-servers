package tokengenerator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

type passwordGenerator struct {
	key []byte
}

func (p *passwordGenerator) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(p.key))
}

func (p *passwordGenerator) generateSecretHash(secret string) (secreteHash string, err error) {
	block, err := aes.NewCipher(p.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	fixedNonce := p.key[:gcm.NonceSize()]
	ciphertext := gcm.Seal(nil, fixedNonce, []byte(secret), nil)
	return hex.EncodeToString(ciphertext), nil
}

func (p *passwordGenerator) parseSecretHash(secreteHash string) (secret string, err error) {
	ciphertext, err := hex.DecodeString(secreteHash)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(p.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("invalid ciphertext")
	}

	fixedNonce := p.key[:gcm.NonceSize()]
	plaintext, err := gcm.Open(nil, fixedNonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func newPasswordGenerator(key []byte) *passwordGenerator {
	return &passwordGenerator{key: key}
}
