package middleware

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/interface/api/context"
	"github.com/education-english-web/BE-english-web/app/services"
	mockServices "github.com/education-english-web/BE-english-web/app/services/mock"
)

type InternalUserAuthenticatorTestSuite struct {
	suite.Suite
	resp *httptest.ResponseRecorder
}

func TestInternalUserAuthenticatorTestSuite(t *testing.T) {
	suite.Run(t, new(InternalUserAuthenticatorTestSuite))
}

func (st *InternalUserAuthenticatorTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	st.resp = httptest.NewRecorder()
}

func (st *InternalUserAuthenticatorTestSuite) TestNewInternalUserAuthenticator() {
	mockCtrl := gomock.NewController(st.T())
	defer mockCtrl.Finish()

	jwt := mockServices.NewMockJWT(mockCtrl)

	type args struct {
		jwt services.JWT
	}

	tests := []struct {
		name string
		args args
		want *UserAuthenticator
	}{
		{
			name: "success",
			args: args{
				jwt: jwt,
			},
			want: &UserAuthenticator{
				jwt: jwt,
			},
		},
	}

	for _, tt := range tests {
		got := NewUserAuthenticator(tt.args.jwt)
		if !st.Equal(tt.want, got) {
			st.T().Errorf("TestNewInternalUserAuthenticator mismatched result:\ngot: %v\nwant: %s", got, tt.want)
		}
	}
}

func (st *InternalUserAuthenticatorTestSuite) TestInternalUserAuthenticator_Authenticate() {
	st.T().Run("jwt token doesnot exist in cookie", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, router := gin.CreateTestContext(st.resp)

		authenticator := &UserAuthenticator{}

		router.GET("/dummy", authenticator.Authenticate, func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/dummy", http.NoBody)

		router.ServeHTTP(recorder, ctx.Request)

		st.Equal(http.StatusUnauthorized, recorder.Code)

		internalUserEmail := context.GetUserEmail(ctx)
		st.Equal("", internalUserEmail)
	})

	st.T().Run("decrypt jwt token failed", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		recorder := httptest.NewRecorder()
		ctx, router := gin.CreateTestContext(st.resp)

		jwt := mockServices.NewMockJWT(mockCtrl)
		jwt.EXPECT().Decrypt("bearer_jwt_token", &entity.UserClaims{}, false).Return(errors.New("error"))

		authenticator := &UserAuthenticator{jwt: jwt}

		router.GET("/dummy", authenticator.Authenticate, func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/dummy", http.NoBody)
		ctx.Request.AddCookie(&http.Cookie{Name: userAccessKey, Value: "bearer_jwt_token"})

		router.ServeHTTP(recorder, ctx.Request)

		st.Equal(http.StatusUnauthorized, recorder.Code)

		internalUserEmail := context.GetUserEmail(ctx)
		st.Equal("", internalUserEmail)
	})

	st.T().Run("jwt token contains empty email", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		recorder := httptest.NewRecorder()
		ctx, router := gin.CreateTestContext(st.resp)

		jwt := mockServices.NewMockJWT(mockCtrl)
		jwt.EXPECT().Decrypt("bearer_jwt_token", &entity.UserClaims{}, false).SetArg(1, entity.UserClaims{
			Email:     "",
			Role:      entity.UserRoleManager,
			CSRFToken: "csrf_token",
		}).Return(nil)

		authenticator := &UserAuthenticator{jwt: jwt}

		router.GET("/dummy", authenticator.Authenticate, func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/dummy", http.NoBody)
		ctx.Request.AddCookie(&http.Cookie{Name: userAccessKey, Value: "bearer_jwt_token"})

		router.ServeHTTP(recorder, ctx.Request)

		st.Equal(http.StatusUnauthorized, recorder.Code)

		internalUserEmail := context.GetUserEmail(ctx)
		st.Equal("", internalUserEmail)
	})

	st.T().Run("jwt token contains invalid csrf_token", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		recorder := httptest.NewRecorder()
		ctx, router := gin.CreateTestContext(st.resp)

		jwt := mockServices.NewMockJWT(mockCtrl)
		jwt.EXPECT().Decrypt("bearer_jwt_token", &entity.UserClaims{}, false).SetArg(1, entity.UserClaims{
			Email:     "email@example.com",
			Role:      entity.UserRoleManager,
			CSRFToken: "csrf_token",
		}).Return(nil)

		authenticator := &UserAuthenticator{jwt: jwt}

		router.POST("/dummy", authenticator.Authenticate, func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/dummy", new(bytes.Buffer))
		ctx.Request.Header.Add(csrfHeader, "csrf_token_1")
		ctx.Request.AddCookie(&http.Cookie{Name: userAccessKey, Value: "bearer_jwt_token"})

		router.ServeHTTP(recorder, ctx.Request)

		st.Equal(http.StatusUnauthorized, recorder.Code)

		internalUserEmail := context.GetUserEmail(ctx)
		st.Equal("", internalUserEmail)
	})

	st.T().Run("success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		internalEmail := "internal@example.com"
		jwt := mockServices.NewMockJWT(mockCtrl)
		jwt.EXPECT().Decrypt("bearer_jwt_token", &entity.UserClaims{}, false).SetArg(1, entity.UserClaims{
			Email:     internalEmail,
			Role:      entity.UserRoleManager,
			CSRFToken: "csrf_token",
		}).Return(nil)

		ctx, _ := gin.CreateTestContext(st.resp)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/testing", new(bytes.Buffer))
		ctx.Request.Header.Add(csrfHeader, "csrf_token")
		ctx.Request.AddCookie(&http.Cookie{Name: userAccessKey, Value: "bearer_jwt_token"})

		authenticator := &UserAuthenticator{jwt}
		authenticator.Authenticate(ctx)

		internalUserEmail := context.GetUserEmail(ctx)
		st.Equal(internalEmail, internalUserEmail)
	})
}

func (st *InternalUserAuthenticatorTestSuite) TestSetInternalUserAccessCookie() {
	st.T().Run("success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx, _ := gin.CreateTestContext(st.resp)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/dummy", new(bytes.Buffer))
		ctx.Request.Header.Add(userAccessKey, "internal_user_access_token")

		SetUserAccessCookie(ctx, "internal_user_access_token")
	})
}
