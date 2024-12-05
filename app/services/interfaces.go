package services

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"github.com/golang-jwt/jwt"
)

// JWT describes what a jwt impl is capable of
type JWT interface {
	Encrypt(claims jwt.Claims) (tokenStr string, err error)
	Decrypt(tokenStr string, claims jwt.Claims, skipClaimsValidation bool) error
}

// HashPass describes what a hash pass impl is capable of
type HashPass interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) bool
}

//type JobService interface {
//	Queue(job entity.Job) error
//	QueueProcessingJobs() error
//	Retry(job entity.Job) error
//}
//
//type SQSService interface {
//	GetMessages(queueName string) error
//}
