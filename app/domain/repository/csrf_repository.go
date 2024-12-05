package repository

import (
	"context"
	"time"
)

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

// CSRFRepository provides interface to handle csrf tokens
type CSRFRepository interface {
	Get(ctx context.Context, userID, uid string) (string, error)
	Save(ctx context.Context, userID, uid, value string, expiresAt time.Duration) error
	DeleteAll(ctx context.Context, userID string) error
}
