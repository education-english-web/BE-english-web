package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserToken struct {
	ID           int       `gorm:"column:id;primaryKey"`
	UserID       uuid.UUID `gorm:"column:user_id"`
	RefreshToken string    `gorm:"column:refresh_token"`
	ExpiredAt    int64     `gorm:"column:expired_at"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
