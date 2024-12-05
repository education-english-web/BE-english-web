package specifications

import (
	"fmt"

	"gorm.io/gorm"
)

type userByEmail struct {
	email string
}

// UserByEmail creates spec for getting user by email
func UserByEmail(email string) I {
	return userByEmail{
		email: email,
	}
}

// GormQuery returns gorm query from conditions
func (q userByEmail) GormQuery(db *gorm.DB) *gorm.DB {
	return db.Table("users").Where("email = ?", q.email)
}

func (q userByEmail) CrucialQueryCondition() string {
	return fmt.Sprintf("email = %q", q.email)
}
