package usecases

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/domain/repository"
	"github.com/education-english-web/BE-english-web/app/domain/repository/specifications"
	"github.com/education-english-web/BE-english-web/pkg/tracer/datadog"
)

type userMeUsecase struct {
	userRepo repository.UserRepository
}

func NewUserMeUsecase(userRepo repository.UserRepository) UserMeUsecase {
	return &userMeUsecase{
		userRepo: userRepo,
	}
}

func (uc *userMeUsecase) Execute(ctx context.Context, userID uuid.UUID) (entity.User, error) {
	span, spanCtx := datadog.StartSpanFromCtx(ctx)
	defer span.Finish()

	user, err := uc.userRepo.Get(spanCtx, specifications.UserByUserID(userID))
	if err != nil {
		return entity.User{}, fmt.Errorf("error while getting internal user by email: %w", err)
	}

	return user, nil
}
