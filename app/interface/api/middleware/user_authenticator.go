package middleware

import (
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/errors"
	"github.com/education-english-web/BE-english-web/app/interface/api/context"
	"github.com/education-english-web/BE-english-web/app/interface/api/presenter"
	"github.com/education-english-web/BE-english-web/app/services"
	"github.com/education-english-web/BE-english-web/pkg/tracer"
)

const (
	userAccessKey = "education_jwt_user"
)

type UserAuthenticator struct {
	jwt services.JWT
}

func NewUserAuthenticator(jwt services.JWT) *UserAuthenticator {
	return &UserAuthenticator{
		jwt: jwt,
	}
}

func (au *UserAuthenticator) Authenticate(ctx *gin.Context) {
	token, err := ctx.Cookie(userAccessKey)
	if err != nil {
		presenter.RenderErrors(ctx, tracer.NoopSpan{}, errors.NewCustomErrorUnauthorized(
			errors.CodeUserUnauthorized,
			"user unauthorized",
			nil,
		))
		ctx.Abort()

		return
	}

	claims := entity.UserClaims{}
	if err := au.jwt.Decrypt(token, &claims, false); err != nil {
		presenter.RenderErrors(ctx, tracer.NoopSpan{}, errors.NewCustomErrorUnauthorized(
			errors.CodeUserUnauthorized,
			"user unauthorized",
			nil,
		))
		ctx.Abort()

		return
	}

	if claims.Email == "" {
		presenter.RenderErrors(ctx, tracer.NoopSpan{}, errors.NewCustomErrorUnauthorized(
			errors.CodeUserUnauthorized,
			"user unauthorized",
			nil,
		))
		ctx.Abort()

		return
	}

	//if claims.ExpiresAt < time.Now().Unix() {
	//	presenter.RenderErrors(ctx, tracer.NoopSpan{}, errors.NewCustomErrorUnauthorized(
	//		errors.CodeUserUnauthorized,
	//		"token expired",
	//		nil,
	//	))
	//	ctx.Abort()
	//	return
	//}

	if !slices.Contains(allowMethods, ctx.Request.Method) {
		csrfToken := ctx.Request.Header.Get(csrfHeader)
		if claims.CSRFToken != csrfToken {
			presenter.RenderErrors(ctx, tracer.NoopSpan{}, errors.NewCustomErrorUnauthorized(
				errors.CodeUserUnauthorized,
				"user unauthorized",
				nil,
			))
			ctx.Abort()

			return
		}
	}

	context.SetUserEmail(ctx, claims.Email)
	context.SetUserID(ctx, claims.UserID)
	context.SetPhoneNumber(ctx, claims.PhoneNumber)
	ctx.Next()
}

// SetUserAccessCookie to set access token for user
func SetUserAccessCookie(ctx *gin.Context, token string) {
	setCookie(ctx, userAccessKey, token)
}
