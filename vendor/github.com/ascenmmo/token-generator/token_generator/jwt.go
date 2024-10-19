package tokengenerator

import (
	"errors"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretHash = []byte("mega me 123 supe secret gangster")

type jwtClaims struct {
	jwt.StandardClaims
	Info tokentype.Info
}

type jwtGenerator struct {
	key []byte
}

func (s *jwtGenerator) generate(info tokentype.Info) (string, error) {
	var claims = jwtClaims{
		Info: info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(info.TTL).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	return token.SignedString(s.key)
}

func (s *jwtGenerator) parse(token string) (info tokentype.Info, err error) {
	claims := jwtClaims{}
	TokenFunction := func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return s.key, nil
	}

	tkn, err := jwt.ParseWithClaims(token, &claims, TokenFunction)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return info, err
		}
		return info, err

	}

	if !tkn.Valid {
		return info, errors.New("invalid jwt token")
	}

	return claims.Info, nil
}

//func (s jwtGeterator) encrypt(hash []byte) ([]byte, error) {
//	block, err := aes.NewCipher(secretHash)
//	if err != nil {
//		return nil, err
//	}
//	b := base64.StdEncoding.EncodeToString(hash)
//	ciphertext := make([]byte, aes.BlockSize+len(b))
//	iv := ciphertext[:aes.BlockSize]
//	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//		return nil, err
//	}
//	cfb := cipher.NewCFBEncrypter(block, iv)
//	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
//	return ciphertext, nil
//}
//
//func (s jwtGeterator) decrypt(hash []byte) ([]byte, error) {
//
//	block, err := aes.NewCipher(secretHash)
//	if err != nil {
//		return nil, err
//	}
//	b := base64.StdEncoding.EncodeToString(hash)
//	ciphertext := make([]byte, aes.BlockSize+len(b))
//	iv := ciphertext[:aes.BlockSize]
//	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//		return nil, err
//	}
//	cfb := cipher.NewCFBEncrypter(block, iv)
//	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
//	return ciphertext, nil
//}

func newJwtGenerator(key []byte) *jwtGenerator {
	return &jwtGenerator{key: key}
}
