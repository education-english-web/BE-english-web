package dto

import "time"

// UserInformation is a struct for user information
type UserInformation struct {
	UserID      string
	UserName    string
	Email       string
	PhoneNumber string
	AvatarLink  string
	RoleCode    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
}

// UserAddRequest is a struct for request
type UserAddRequest struct {
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
	AvatarLink  string
	RoleCode    string
}

// UserLoginRequest is a struct for request
type UserLoginRequest struct {
	Email    string
	Phone    string
	Password string
}

type UserLoginResponse struct {
	UserID string `json:"user_id"`

	UserName string `json:"username"`
	Email    string `json:"email"`

	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	CSRFToken string `json:"csrf_token"`
}

// UserRefreshTokenRequest is a struct for request
type UserRefreshTokenRequest struct {
	RefreshToken string
}
