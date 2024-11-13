package tokengenerator

import (
	"errors"
	"fmt"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"strings"
)

type TokenGenerator interface {
	GenerateToken(info tokentype.Info, tokenType string) (token string, err error)
	ParseToken(tokenWithType string) (info tokentype.Info, err error)

	GenerateUniqueHash(txt string) (hashData string, err error)
	ParseUniqueHash(hashData string) (txt string, err error)

	PasswordHash(password string) (passwordHash string)

	GenerateSecretHash(secret string) (secretHash string, err error)
	ParseSecretHash(secretHash string) (secret string, err error)
}

const (
	AESGCM = "a"
	JWT    = "j"
)

type tokenGenerator struct {
	aesgcm            *aesgmGenerator
	jwt               *jwtGenerator
	hashGenerator     *hashGenerator
	passwordGenerator *passwordGenerator

	key []byte
}

func (t *tokenGenerator) GenerateToken(info tokentype.Info, tokenType string) (token string, err error) {
	info.SetDefaultValueIfEmpty()

	switch tokenType {
	case AESGCM:
		generate, err := t.aesgcm.generate(info)
		if err != nil {
			return "", err
		}
		withType := setTokenAndType(generate, tokenType)

		return withType, err
	case JWT:
		generate, err := t.jwt.generate(info)
		if err != nil {
			return "", err
		}
		withType := setTokenAndType(generate, tokenType)

		return withType, err
	default:
		return "", errors.New("invalid token type")
	}
}

func (t *tokenGenerator) ParseToken(tokenWithType string) (info tokentype.Info, err error) {
	tokenType, token := getTokenAndType(tokenWithType)
	switch tokenType {
	case AESGCM:
		return t.aesgcm.parse(token)
	case JWT:
		return t.jwt.parse(token)
	default:
		return info, errors.New("invalid token type")
	}
}

func (t *tokenGenerator) GenerateUniqueHash(txt string) (hashData string, err error) {
	return t.hashGenerator.generateHash(txt)
}

func (t *tokenGenerator) ParseUniqueHash(hashData string) (txt string, err error) {
	return t.hashGenerator.parseHash(hashData)
}

func (t *tokenGenerator) PasswordHash(password string) (passwordHash string) {
	return t.passwordGenerator.generatePasswordHash(password)
}

func (t *tokenGenerator) GenerateSecretHash(secret string) (secretHash string, err error) {
	return t.passwordGenerator.generateSecretHash(secret)
}

func (t *tokenGenerator) ParseSecretHash(secretHash string) (secret string, err error) {
	return t.passwordGenerator.parseSecretHash(secretHash)
}

func getTokenAndType(tokenWithType string) (tokenType string, token string) {
	split := strings.Split(tokenWithType, "$_$")
	if len(split) != 2 {
		return
	}
	token = split[0]
	tokenType = split[1]
	return
}

func setTokenAndType(token string, tokenType string) (tokenWithType string) {
	tokenWithType = fmt.Sprintf("%s$_$%s", token, tokenType)
	return tokenWithType
}

func NewTokenGenerator(key string) (TokenGenerator, error) {
	buf := []byte(key)
	if len(buf) != 32 {
		return nil, errors.New("invalid key length mast be 32 bytes")
	}

	return &tokenGenerator{
		key:               buf,
		jwt:               newJwtGenerator(buf),
		aesgcm:            newAESGMGenerator(buf),
		passwordGenerator: newPasswordGenerator(buf),
		hashGenerator:     newHashGenerator(buf),
	}, nil
}
