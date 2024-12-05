package repository

import (
	"context"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/domain/repository/specifications"
)

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

// UserRepository provides the way to interact with user data
type UserTokenRepository interface {
	Create(ctx context.Context, userToken *entity.UserToken) error
	Get(ctx context.Context, specs specifications.I) (entity.UserToken, error)
	Update(ctx context.Context, userToken *entity.UserToken) error
	//Delete(ctx context.Context, specs specifications.I) error
}
