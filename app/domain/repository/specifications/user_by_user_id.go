package specifications

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userByUserID struct {
	userID uuid.UUID
}

// UserByUserID creates spec for getting user by user id
func UserByUserID(userID uuid.UUID) I {
	return userByUserID{
		userID: userID,
	}
}

// GormQuery returns gorm query from conditions
func (q userByUserID) GormQuery(db *gorm.DB) *gorm.DB {
	return db.Table("users").Where("user_id = ?", q.userID)
}

func (q userByUserID) CrucialQueryCondition() string {
	return fmt.Sprintf("user_id = %q", q.userID)
}
