package services

import (
	"golang.org/x/crypto/bcrypt"
)

type hashPass struct {
	hash string
}

func NewHashPass(salt string) HashPass {
	return &hashPass{
		hash: salt,
	}
}

func (hp *hashPass) HashPassword(password string) (string, error) {
	salt := hp.hash

	saltedPassword := password + salt

	hash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (hp *hashPass) VerifyPassword(password, hash string) bool {
	salt := hp.hash
	saltedPassword := password + salt

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedPassword))
	return err == nil
}
