package usecases

import (
	"context"
	"fmt"

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

type userRefreshTokenUsecase struct {
	userRepo      repository.UserRepository
	userTokenRepo repository.UserTokenRepository
	timeFactory   timeutil.TimeFactory
	jwt           services.JWT
	uuid          uuidstring.UUIDString
}

func NewUserRefreshTokenUsecase(
	userRepo repository.UserRepository,
	userTokenRepo repository.UserTokenRepository,
	timeFactory timeutil.TimeFactory,
	jwt services.JWT,
	uuid uuidstring.UUIDString,
) RefreshTokenUsecase {
	return &userRefreshTokenUsecase{
		userRepo:      userRepo,
		userTokenRepo: userTokenRepo,
		timeFactory:   timeFactory,
		jwt:           jwt,
		uuid:          uuid,
	}
}

func (uc *userRefreshTokenUsecase) Execute(ctx context.Context, dtoReq dto.UserRefreshTokenRequest) (dto.UserLoginResponse, error) {
	span, spanCtx := datadog.StartSpanFromCtx(ctx)
	defer span.Finish()

	checkClaimsRefresh := entity.RefreshTokenClaims{}
	checkExpRefreshToken := uc.jwt.Decrypt(dtoReq.RefreshToken, &checkClaimsRefresh, false)
	if checkExpRefreshToken != nil {
		return dto.UserLoginResponse{}, appErrors.NewErrorInvalidArgument(appErrors.CodeNotFound, "token validation failed. Please login again", nil)
	}

	user, err := uc.userRepo.Get(spanCtx, specifications.UserTokenByRefreshToken(checkClaimsRefresh.UserID, dtoReq.RefreshToken))
	if err != nil {
		return dto.UserLoginResponse{}, appErrors.NewErrorInvalidArgument(appErrors.CodeNotFound, "token not found. Please login again or contact admin", nil)
	}

	user, err = uc.userRepo.Get(spanCtx, specifications.UserByUserID(user.UserID))
	if err != nil {
		return dto.UserLoginResponse{}, appErrors.NewErrorInvalidArgument(appErrors.CodeNotFound, "user not found", nil)
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
		return dto.UserLoginResponse{}, fmt.Errorf("error while encrypting internal user claims %w for refresh token", err)
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
