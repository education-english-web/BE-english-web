package repository

import (
	"context"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/domain/repository/specifications"
)

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

// UserRepository provides the way to interact with user data
type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	//Update(ctx context.Context, internalUser *entity.InternalUser) error
	Get(ctx context.Context, specs specifications.I) (entity.User, error)
	//GetAll(ctx context.Context, paging entity.PagingRequest) ([]entity.InternalUser, error)
	//Count(ctx context.Context, spec specifications.I) (uint32, error)
}
