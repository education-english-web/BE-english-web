package service

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"golang.org/x/oauth2"
)

// OIDCService represents the inteface with OP
type OIDCService interface {
	GetToken(code, nonce, redirectURI string) (oauth2.Token, error)
	GetClaims(token oauth2.Token, claims interface{}) error
	RefreshToken(token oauth2.Token) (oauth2.Token, error)
	GetUserInfo(token oauth2.Token) (OAuth2UserInfo, error)
	GetTokenFromJwtBearer(username string) (oauth2.Token, error)
}

type OAuth2UserInfo struct {
	Subject       string
	Profile       string
	Email         string
	EmailVerified bool

	Claims map[string]interface{}
}
