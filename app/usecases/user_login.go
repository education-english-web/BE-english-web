package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/domain/repository"
	"github.com/education-english-web/BE-english-web/app/domain/repository/specifications"
	appErrors "github.com/education-english-web/BE-english-web/app/errors"
	"github.com/education-english-web/BE-english-web/app/services"
	"github.com/education-english-web/BE-english-web/app/usecases/dto"
	"github.com/education-english-web/BE-english-web/pkg/timeutil"
	"github.com/education-english-web/BE-english-web/pkg/tracer/datadog"
	"github.com/education-english-web/BE-english-web/pkg/uuidstring"
)

var (
	userValidityPeriodAccessToken  = time.Hour * 12
	userValidityPeriodRefreshToken = time.Hour * 24 * 30 // 30 days
)

type userLoginUsecase struct {
	userRepo      repository.UserRepository
	userTokenRepo repository.UserTokenRepository
	timeFactory   timeutil.TimeFactory
	hashPass      services.HashPass
	jwt           services.JWT
	uuid          uuidstring.UUIDString
}

func NewUserLoginUsecase(
	userRepo repository.UserRepository,
	userTokenRepo repository.UserTokenRepository,
	timeFactory timeutil.TimeFactory,
	uuid uuidstring.UUIDString,
	hashPass services.HashPass,
	jwt services.JWT,
) UserLoginUsecase {
	return &userLoginUsecase{
		userRepo:      userRepo,
		userTokenRepo: userTokenRepo,
		timeFactory:   timeFactory,
		uuid:          uuid,
		hashPass:      hashPass,
		jwt:           jwt,
	}
}

// Execute token for the user
func (uc *userLoginUsecase) Execute(ctx context.Context, dtoReq dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	span, spanCtx := datadog.StartSpanFromCtx(ctx)
	defer span.Finish()

	user, err := uc.userRepo.Get(spanCtx, specifications.UserByEmail(dtoReq.Email))
	if err != nil {
		return dto.UserLoginResponse{}, appErrors.NewErrorInvalidArgument(appErrors.CodeNotFound, "user not found", dtoReq.Email)
	}

	if user.IsDeleted == true {
		return dto.UserLoginResponse{}, appErrors.NewErrorInvalidArgument(appErrors.CodeInternalUserExisted, "user is deleted. Please contact admin", dtoReq.Email)
	}

	if !uc.hashPass.VerifyPassword(dtoReq.Password, user.Password) {
		return dto.UserLoginResponse{}, appErrors.NewErrorInvalidArgument(appErrors.CodeInvalidPassword, "invalid password", dtoReq.Email)
	}

	// grant access
	now := uc.timeFactory.Now()
	csrfToken := uc.uuid.GetUUID()
	issuedAt := now
	expiresAt := issuedAt.Add(userValidityPeriodAccessToken)
	claims := entity.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  issuedAt.Unix(),
		},
		UserID:      user.UserID,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.RoleCode,
		CSRFToken:   csrfToken,
	}

	accessToken, err := uc.jwt.Encrypt(claims)
	if err != nil {
		return dto.UserLoginResponse{}, fmt.Errorf("error while encrypting user claims %w for access token", err)
	}

	expiresAtRefreshToken := issuedAt.Add(userValidityPeriodRefreshToken)
	claimsRefresh := entity.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAtRefreshToken.Unix(),
		},
		UserID: user.UserID,
		Scope:  user.RoleCode.String(),
	}
	refreshToken, err := uc.jwt.Encrypt(claimsRefresh)
	if err != nil {
		return dto.UserLoginResponse{}, fmt.Errorf("error while encrypting internal user claims %w for reefresh token", err)
	}

	userToken, err := uc.userTokenRepo.Get(spanCtx, specifications.UserTokenByUserID(user.UserID))
	if err != nil {
		if !errors.Is(err, appErrors.ErrNotFound) {
			return dto.UserLoginResponse{}, fmt.Errorf("error while getting user token by user id: %w", err)
		}
	}

	if userToken == (entity.UserToken{}) {
		userToken = entity.UserToken{
			UserID:       user.UserID,
			RefreshToken: refreshToken,
			ExpiredAt:    claimsRefresh.StandardClaims.ExpiresAt,
			CreatedAt:    now,
		}
		err = uc.userTokenRepo.Create(spanCtx, &userToken)
		if err != nil {
			return dto.UserLoginResponse{}, fmt.Errorf("error while creating user token: %w", err)
		}
	} else {
		userToken.RefreshToken = refreshToken
		userToken.ExpiredAt = claimsRefresh.StandardClaims.ExpiresAt
		userToken.UpdatedAt = now
		err = uc.userTokenRepo.Update(spanCtx, &userToken)
		if err != nil {
			return dto.UserLoginResponse{}, fmt.Errorf("error while updating user token: %w", err)
		}
	}

	return dto.UserLoginResponse{
		UserID:       user.UserID.String(),
		UserName:     user.UserName,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CSRFToken:    csrfToken,
	}, nil
}
