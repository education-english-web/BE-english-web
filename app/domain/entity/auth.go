package entity

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// AuthClaims the claim for authentication
type AuthClaims struct {
	jwt.StandardClaims

	UserName          string    `json:"username,omitempty"`
	UserID            uuid.UUID `json:"userId,omitempty"`
	Exp               int64     `json:"exp,omitempty"`
	Roles             int       `json:"roles"`
	AcceptedTerms     bool      `json:"accepted_terms,omitempty"`
	ProxyLoginEventID *uint32   `json:"proxy_login_event_id,omitempty"`
	CSRFToken         string    `json:"csrf_token,omitempty"`
}
