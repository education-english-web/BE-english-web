package usecases

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/domain/repository"
	"github.com/education-english-web/BE-english-web/app/domain/repository/specifications"
	appErrors "github.com/education-english-web/BE-english-web/app/errors"
	"github.com/education-english-web/BE-english-web/app/services"
	"github.com/education-english-web/BE-english-web/app/usecases/dto"
	"github.com/education-english-web/BE-english-web/pkg/timeutil"
	"github.com/education-english-web/BE-english-web/pkg/tracer/datadog"
)

type userAddUsecase struct {
	userRepo     repository.UserRepository
	timeFactory  timeutil.TimeFactory
	hashPassword services.HashPass
}

func NewUserAddUsecase(
	userRepo repository.UserRepository,
	timeFactory timeutil.TimeFactory,
	hashPassword services.HashPass,
) UserAddUsecase {
	return &userAddUsecase{
		userRepo:     userRepo,
		timeFactory:  timeFactory,
		hashPassword: hashPassword,
	}
}

// Execute adds an user to the system
func (uc *userAddUsecase) Execute(ctx context.Context, dtoReq dto.UserAddRequest) (entity.User, error) {
	span, spanCtx := datadog.StartSpanFromCtx(ctx)
	defer span.Finish()

	newUser, err := uc.userRepo.Get(spanCtx, specifications.UserByEmail(dtoReq.Email))
	if err == nil {
		return entity.User{}, appErrors.NewErrorInvalidArgument(appErrors.CodeInternalUserExisted, "same user already exists", dtoReq.Email)
	}

	if newUser.IsDeleted == true {
		return entity.User{}, appErrors.NewErrorInvalidArgument(appErrors.CodeInternalUserExisted, "user is deleted. Please contact admin", dtoReq.Email)
	}

	if !errors.Is(err, appErrors.ErrNotFound) {
		return entity.User{}, fmt.Errorf("error while getting internal user by email: %w", err)
	}

	passwordHash, _ := uc.hashPassword.HashPassword(dtoReq.Password)

	newUser = entity.User{
		UserID:      uuid.New(),
		Email:       dtoReq.Email,
		RoleCode:    entity.GetUserRole(dtoReq.RoleCode),
		UserName:    dtoReq.UserName,
		Password:    passwordHash,
		PhoneNumber: dtoReq.PhoneNumber,
		AvatarLink:  dtoReq.AvatarLink,
	}

	if err = uc.userRepo.Create(spanCtx, &newUser); err != nil {
		return entity.User{}, fmt.Errorf("error while creating user: %w", err)
	}

	return newUser, nil
}
