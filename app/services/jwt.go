package services

import (
	"github.com/golang-jwt/jwt"
)

// jwtGo hold secrets (key, public/private key) for encrypt & decrypt
type jwtGo struct {
	jwtSecret []byte
}

// NewJWT returns JWT instance for encrypt/decrypt jwt tokens
func NewJWT(jwtSecret string) JWT {
	return &jwtGo{
		jwtSecret: []byte(jwtSecret),
	}
}

// Encrypt encrypts jwt claims
func (uc *jwtGo) Encrypt(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(uc.jwtSecret)

	return tokenStr, err
}

// Decrypt decrypts and verifys jwt claims
// warning: please pass skipClaimsValidation as FALSE by default UNLESS you know what you are doing
func (uc *jwtGo) Decrypt(tokenStr string, claims jwt.Claims, skipClaimsValidation bool) error {
	parser := jwt.Parser{
		SkipClaimsValidation: skipClaimsValidation,
	}

	if _, err := parser.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) { return uc.jwtSecret, nil },
	); err != nil {
		return err
	}

	return nil
}
