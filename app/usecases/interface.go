package usecases

import (
	"context"

	"github.com/google/uuid"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/usecases/dto"
)

type UserMeUsecase interface {
	Execute(ctx context.Context, userID uuid.UUID) (entity.User, error)
}

// UserAddUsecase provides a method to add an internal user to the system
type UserAddUsecase interface {
	Execute(ctx context.Context, dto dto.UserAddRequest) (entity.User, error)
}

// UserLoginUsecase provides a method to login an user to the system
type UserLoginUsecase interface {
	Execute(ctx context.Context, dto dto.UserLoginRequest) (dto.UserLoginResponse, error)
}

// RefreshTokenUsecase provides a method to refresh the token
type RefreshTokenUsecase interface {
	Execute(ctx context.Context, dto dto.UserRefreshTokenRequest) (dto.UserLoginResponse, error)
}
