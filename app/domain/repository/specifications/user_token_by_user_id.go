package specifications

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userTokenByUserID struct {
	userID uuid.UUID
}

// UserTokenByUserID creates spec for getting user by email
func UserTokenByUserID(userID uuid.UUID) I {
	return userTokenByUserID{
		userID: userID,
	}
}

// GormQuery returns gorm query from conditions
func (q userTokenByUserID) GormQuery(db *gorm.DB) *gorm.DB {
	return db.Table("user_tokens").Where("user_id = ?", q.userID)
}

func (q userTokenByUserID) CrucialQueryCondition() string {
	return fmt.Sprintf("user_id = %q", q.userID)
}
