package tokengenerator

import (
	"crypto/sha256"
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

func newPasswordGenerator(key []byte) *passwordGenerator {
	return &passwordGenerator{key: key}
}
