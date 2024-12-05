package specifications

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//UserByRefreshToken

type userTokenByRefreshToken struct {
	userID       uuid.UUID
	refreshToken string
}

// UserTokenByRefreshToken creates spec for getting user by email
func UserTokenByRefreshToken(userID uuid.UUID, refreshToken string) I {
	return userTokenByRefreshToken{
		userID:       userID,
		refreshToken: refreshToken,
	}
}

// GormQuery returns gorm query from conditions
func (q userTokenByRefreshToken) GormQuery(db *gorm.DB) *gorm.DB {
	return db.Table("user_tokens").Where("refresh_token = ? and user_id = ?", q.refreshToken, q.userID)
}

func (q userTokenByRefreshToken) CrucialQueryCondition() string {
	return fmt.Sprintf("refresh_token = %q and user_id = %q", q.refreshToken, q.userID)
}
