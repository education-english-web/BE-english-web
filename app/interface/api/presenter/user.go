package presenter

import (
	"github.com/education-english-web/BE-english-web/app/domain/entity"
)

// InternalUser holds response information of internal users
type User struct {
	UserID      string `json:"user_id"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	AvatarLink  string `json:"avatar_link"`
	RoleCode    string `json:"role"`
}

// FormInternalUser convert an internal user entity struct to its presenter struct
func FormInternalUser(internalUser entity.User) User {
	return User{
		UserID:      internalUser.UserID.String(),
		Email:       internalUser.Email,
		UserName:    internalUser.UserName,
		PhoneNumber: internalUser.PhoneNumber,
		AvatarLink:  internalUser.AvatarLink,
		RoleCode:    internalUser.RoleCode.String(),
	}
}

// FormInternalUsers converts a list of internal user entity struct to a list of its presenter struct
func FormInternalUsers(internalUsers []entity.User) []User {
	presenter := make([]User, len(internalUsers))
	for i, user := range internalUsers {
		presenter[i] = FormInternalUser(user)
	}

	return presenter
}
