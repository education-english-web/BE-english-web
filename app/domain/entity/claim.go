package entity

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserClaims struct {
	jwt.StandardClaims
	UserID      uuid.UUID `json:"user_id,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Role        UserRole  `json:"role,omitempty"`
	CSRFToken   string    `json:"csrf_token,omitempty"`
}

// HandoverSessionClaims contains information related to a handover session
type HandoverSessionClaims struct {
	jwt.StandardClaims
	TokenID string `json:"token_id,omitempty"`
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"user_id,omitempty"`
	Scope  string    `json:"scope,omitempty"`
}
