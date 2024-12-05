package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/education-english-web/BE-english-web/app/errors"
	"github.com/education-english-web/BE-english-web/app/interface/api/middleware"
	"github.com/education-english-web/BE-english-web/app/interface/api/payload"
	"github.com/education-english-web/BE-english-web/app/interface/api/presenter"
	"github.com/education-english-web/BE-english-web/app/registry"
	"github.com/education-english-web/BE-english-web/app/usecases/dto"
	"github.com/education-english-web/BE-english-web/pkg/tracer/datadog"
)

// UserHandler provides the way to login for aweb users
type UserHandler struct {
	BaseHandler
}

// NewUserHandler generate UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Add adds an user to system
// @Summary add an internal user to system
// @Description add an user to system
// @Tags users
// @Accept  json
// @Produce json
// @Param   payloadBody body payload.UserAddRequest true "internal user email and role"
// @Success 200 {object} presenter.User
// @Router /users/sign-up [post]
func (hdl *UserHandler) Add(ctx *gin.Context) {
	span, spanCtx := datadog.StartSpanFromCtx(ctx.Request.Context())
	defer span.Finish()

	var payl payload.UserAddRequest
	if err := ctx.ShouldBindJSON(&payl); err != nil {
		presenter.RenderErrors(ctx, span, errors.NewErrorInvalidArgument(
			errors.CodeInvalidPayload,
			"invalid payload json",
			"",
		))

		return
	}

	if err := payl.Validate(); err != nil {
		presenter.RenderErrors(ctx, span, err)

		return
	}

	internalUser, err := registry.InjectedUserAddUsecase().Execute(spanCtx, dto.UserAddRequest{
		UserName:    *payl.UserName,
		Email:       *payl.Email,
		Password:    *payl.Password,
		PhoneNumber: *payl.PhoneNumber,
		RoleCode:    *payl.RoleCode,
	})
	if err != nil {
		presenter.RenderErrors(ctx, span, err)

		return
	}

	presenter.RenderData(ctx, presenter.FormInternalUser(internalUser), nil)
}

// Login login
// @Summary login an user
// @Description login an user
// @Tags users
// @Accept  json
// @Produce json
// @Param   payloadBody body payload.UserLoginRequest true "internal user email and password"
// @Success 200 {object} presenter.User
// @Router /users/login [post]
func (hdl *UserHandler) Login(ctx *gin.Context) {
	span, spanCtx := datadog.StartSpanFromCtx(ctx.Request.Context())
	defer span.Finish()

	var payl payload.UserLoginRequest
	if err := ctx.ShouldBindJSON(&payl); err != nil {
		presenter.RenderErrors(ctx, span, errors.NewErrorInvalidArgument(
			errors.CodeInvalidPayload,
			"invalid payload json",
			"",
		))

		return
	}

	if err := payl.Validate(); err != nil {
		presenter.RenderErrors(ctx, span, err)

		return
	}

	resp, err := registry.InjectedUserLoginUsecase().Execute(spanCtx, dto.UserLoginRequest{
		Email:    *payl.Email,
		Password: *payl.Password,
	})
	if err != nil {
		presenter.RenderErrors(ctx, span, err)

		return
	}

	middleware.SetUserAccessCookie(ctx, resp.AccessToken)
	presenter.RenderData(ctx, resp, nil)
}

// Me function in handler to return logged-in user information
// @Summary Get information of logged in users
// @Description Get information of logged in users
// @Tags users
// @Accept  json
// @Produce json
// @Success 200 {object} presenter.User
// @Router /users/me [get]
func (hdl *UserHandler) Me(ctx *gin.Context) {
	span, spanCtx := datadog.StartSpanFromCtx(ctx.Request.Context())
	defer span.Finish()

	userIDStr := ctx.Query("user_id")
	if userIDStr == "" {
		presenter.RenderErrors(ctx, span, errors.NewErrorInvalidArgument(
			errors.CodeBadRequest,
			"user_id is required",
			"",
		))
		return
	}

	// convert string to uuid
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		presenter.RenderErrors(ctx, span, errors.NewErrorInvalidArgument(
			errors.CodeBadRequest,
			"invalid user_id",
			"",
		))

		return
	}

	userIDCtx, err := hdl.GetUserID(ctx)
	if err != nil {
		presenter.RenderErrors(ctx, span, err)

		return
	}

	if userIDCtx != userID {
		presenter.RenderErrors(ctx, span, errors.NewErrorInvalidArgument(
			errors.CodeUnauthorized,
			"unauthorized",
			"",
		))

		return
	}

	user, err := registry.InjectedUserMeUsecase().Execute(spanCtx, userID)
	if err != nil {
		presenter.RenderErrors(ctx, span, err)

		return
	}

	presenter.RenderData(ctx, presenter.FormInternalUser(user), nil)
}

// RefreshToken refresh token
// @Summary refresh for access token
// @Description refresh  for access token
// @Tags users
// @Accept  json
// @Produce json
// @Success 200 {object} presenter.User
// @Router /users/refresh [post]
func (hdl *UserHandler) RefreshToken(ctx *gin.Context) {
	span, spanCtx := datadog.StartSpanFromCtx(ctx.Request.Context())
	defer span.Finish()

	var payl payload.UserRefreshTokenRequest
	if err := ctx.ShouldBindJSON(&payl); err != nil {
		presenter.RenderErrors(ctx, span, errors.NewErrorInvalidArgument(
			errors.CodeInvalidPayload,
			"invalid payload json",
			"",
		))

		return
	}

	if err := payl.Validate(); err != nil {
		presenter.RenderErrors(ctx, span, err)

		return
	}

	resp, err := registry.InjectedUserRefreshTokenUsecase().Execute(spanCtx, dto.UserRefreshTokenRequest{
		RefreshToken: *payl.RefreshToken,
	})
	if err != nil {
		presenter.RenderErrors(ctx, span, err)

		return
	}

	middleware.SetUserAccessCookie(ctx, resp.AccessToken)
	presenter.RenderData(ctx, resp, nil)
}
