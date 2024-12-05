package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	apiContext "github.com/education-english-web/BE-english-web/app/interface/api/context"
	"github.com/education-english-web/BE-english-web/app/services"
	mockServices "github.com/education-english-web/BE-english-web/app/services/mock"
)

type CookieAuthenticatorTestSuite struct {
	suite.Suite
	resp *httptest.ResponseRecorder
}

func TestCookieAuthenticatorTestSuite(t *testing.T) {
	suite.Run(t, new(CookieAuthenticatorTestSuite))
}

func (st *CookieAuthenticatorTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	st.resp = httptest.NewRecorder()
}

func (st *CookieAuthenticatorTestSuite) TestNewCookieAuthenticator() {
	mockCtrl := gomock.NewController(st.T())
	defer mockCtrl.Finish()

	jwt := mockServices.NewMockJWT(mockCtrl)
	validators := []ClaimsValidator{
		NewUserValidator(),
		NewRolesValidator(),
		NewTermsOfUseValidator(),
	}

	type args struct {
		jwt        services.JWT
		validators []ClaimsValidator
	}

	tests := []struct {
		name string
		args args
		want CookieAuthenticator
	}{
		{
			name: "success",
			args: args{
				jwt:        jwt,
				validators: validators,
			},
			want: CookieAuthenticator{
				jwt:        jwt,
				validators: validators,
			},
		},
	}

	for _, tt := range tests {
		st.T().Run(tt.name, func(t *testing.T) {
			got := NewCookieAuthenticator(tt.args.jwt, tt.args.validators...)

			assert.Equal(t, tt.want, got)
		})
	}
}

func (st *CookieAuthenticatorTestSuite) TestAuthenticateAccess() {
	st.T().Parallel()

	st.T().Run("access key does not exist", func(t *testing.T) {
		ctx, _ := gin.CreateTestContext(st.resp)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/testing", http.NoBody)

		m := CookieAuthenticator{}
		m.Authenticate(ctx)

		userID := apiContext.GetUserID(ctx)
		st.Equal(uint32(0), userID)
		officeID := apiContext.GetNavisOfficeID(ctx)
		st.Equal(uint32(0), officeID)
		mfidUserID := apiContext.GetMFIDUserID(ctx)
		st.Equal(uint32(0), mfidUserID)
		proxyLoginEventID := apiContext.GetProxyLoginEventID(ctx)
		st.Equal(uint32(0), proxyLoginEventID)
	})

	st.T().Run("decrypt access token failed", func(t *testing.T) {
		mockCtrl := gomock.NewController(st.T())
		defer mockCtrl.Finish()

		jwt := mockServices.NewMockJWT(mockCtrl)

		ctx, _ := gin.CreateTestContext(st.resp)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/testing", http.NoBody)
		ctx.Request.AddCookie(&http.Cookie{Name: AccessKey, Value: "access_token"})

		jwt.EXPECT().Decrypt("access_token", &entity.AuthClaims{}, false).Return(errors.New("decrypt failed"))

		m := CookieAuthenticator{
			jwt: jwt,
		}
		m.Authenticate(ctx)

		userID := apiContext.GetUserID(ctx)
		st.Equal(uint32(0), userID)
		officeID := apiContext.GetNavisOfficeID(ctx)
		st.Equal(uint32(0), officeID)
		mfidUserID := apiContext.GetMFIDUserID(ctx)
		st.Equal(uint32(0), mfidUserID)
		proxyLoginEventID := apiContext.GetProxyLoginEventID(ctx)
		st.Equal(uint32(0), proxyLoginEventID)
	})

	st.T().Run("stop by one of validators", func(t *testing.T) {
		mockCtrl := gomock.NewController(st.T())
		defer mockCtrl.Finish()

		jwt := mockServices.NewMockJWT(mockCtrl)

		ctx, _ := gin.CreateTestContext(st.resp)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/testing", http.NoBody)
		ctx.Request.AddCookie(&http.Cookie{Name: AccessKey, Value: "access_token"})

		jwt.EXPECT().Decrypt("access_token", &entity.AuthClaims{}, false).SetArg(1, entity.AuthClaims{
			//MFIDUserID: 1,
			UserID: uuid.UUID{1},
			//TenantUID:  1111,
		}).Return(nil)

		m := CookieAuthenticator{
			jwt: jwt,
			validators: []ClaimsValidator{
				NewUserValidator(),
			},
		}
		m.Authenticate(ctx)

		userID := apiContext.GetUserID(ctx)
		st.Equal(uint32(0), userID)
		officeID := apiContext.GetNavisOfficeID(ctx)
		st.Equal(uint32(0), officeID)
		mfidUserID := apiContext.GetMFIDUserID(ctx)
		st.Equal(uint32(0), mfidUserID)
		proxyLoginEventID := apiContext.GetProxyLoginEventID(ctx)
		st.Equal(uint32(0), proxyLoginEventID)
	})

	st.T().Run("success", func(t *testing.T) {
		mockCtrl := gomock.NewController(st.T())
		defer mockCtrl.Finish()

		jwt := mockServices.NewMockJWT(mockCtrl)

		ctx, _ := gin.CreateTestContext(st.resp)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/testing", http.NoBody)
		ctx.Request.AddCookie(&http.Cookie{Name: AccessKey, Value: "access_token"})

		jwt.EXPECT().Decrypt("access_token", &entity.AuthClaims{}, false).SetArg(1, entity.AuthClaims{
			//MFIDUserID: 1,
			UserID: uuid.UUID{1},
			//OfficeID:   111,
			//TenantUID:  1111,
			ProxyLoginEventID: func() *uint32 {
				i := uint32(1111)

				return &i
			}(),
		}).Return(nil)

		m := CookieAuthenticator{
			jwt: jwt,
			validators: []ClaimsValidator{
				NewUserValidator(),
				NewRolesValidator(),
			},
		}
		m.Authenticate(ctx)

		userID := apiContext.GetUserID(ctx)
		st.Equal(uint32(11), userID)
		officeID := apiContext.GetNavisOfficeID(ctx)
		st.Equal(uint32(111), officeID)
		mfidUserID := apiContext.GetMFIDUserID(ctx)
		st.Equal(uint32(1), mfidUserID)
		proxyLoginEventID := apiContext.GetProxyLoginEventID(ctx)
		st.Equal(uint32(1111), proxyLoginEventID)
	})
}
