package middleware

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	mockRepository "github.com/education-english-web/BE-english-web/app/domain/repository/mock"
	appErrors "github.com/education-english-web/BE-english-web/app/errors"
)

func TestNewUserValidator(t *testing.T) {
	tests := []struct {
		name string
		want ClaimsValidator
	}{
		{
			name: "success",
			want: &userValidator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserValidator()

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_userValidator_validate(t *testing.T) {
	type args struct {
		ctx   *gin.Context
		claim entity.AuthClaims
	}

	tests := []struct {
		name    string
		v       *userValidator
		args    args
		wantErr error
	}{
		{
			name: "user_id is zero",
			v:    &userValidator{},
			args: args{
				ctx: &gin.Context{},
				claim: entity.AuthClaims{
					UserID: 0,
				},
			},
			wantErr: appErrors.NewErrorUnauthorized(),
		},
		{
			name: "user_id is positive",
			v:    &userValidator{},
			args: args{
				ctx: &gin.Context{},
				claim: entity.AuthClaims{
					UserID: 1,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &userValidator{}
			gotErr := v.validate(tt.args.ctx, tt.args.claim)

			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}

func TestNewOfficeValidator(t *testing.T) {
	tests := []struct {
		name string
		want ClaimsValidator
	}{
		{
			name: "success",
			want: &officeValidator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOfficeValidator()

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_officeValidator_validate(t *testing.T) {
	type args struct {
		ctx    *gin.Context
		claims entity.AuthClaims
	}

	tests := []struct {
		name    string
		v       *officeValidator
		args    args
		wantErr error
	}{
		{
			name: "office_id is zero",
			v:    &officeValidator{},
			args: args{
				ctx: &gin.Context{},
				claims: entity.AuthClaims{
					OfficeID: 0,
				},
			},
			wantErr: appErrors.NewErrorForbidden(),
		},
		{
			name: "office_id is positive",
			v:    &officeValidator{},
			args: args{
				ctx: &gin.Context{},
				claims: entity.AuthClaims{
					OfficeID: 1,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &officeValidator{}
			gotErr := v.validate(tt.args.ctx, tt.args.claims)

			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}

func TestNewCSRFValidator(t *testing.T) {
	tests := []struct {
		name string
		want ClaimsValidator
	}{
		{
			name: "success",
			want: &csrfValidator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCSRFValidator(nil)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_csrfValidator_validate(t *testing.T) {
	t.Parallel()

	t.Run("failed to get stored csrf token", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mfidUserID := uint32(1)
		csrfToken := "csrf_token"

		ctx := &gin.Context{
			Request: &http.Request{
				Method: http.MethodPost,
				URL: &url.URL{
					Path: "/api/v1/contracts/abcxyz",
				},
			},
		}

		csrfRepository := mockRepository.NewMockCSRFRepository(mockCtrl)
		csrfRepository.EXPECT().Get(ctx, mfidUserID, csrfToken).Return("", errors.New("zzz"))

		v := csrfValidator{
			csrfRepository: csrfRepository,
		}

		wantErr := appErrors.NewErrorUnauthorized()
		gotErr := v.validate(ctx, entity.AuthClaims{
			MFIDUserID: mfidUserID,
			CSRFToken:  csrfToken,
		})

		assert.Equal(t, wantErr, gotErr)
	})

	//#nosec
	t.Run("stored csrf token is different with token from header", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mfidUserID := uint32(1)
		csrfToken := "csrf_token"
		csrfTokenStored := "csrf_token_stored"

		ctx := &gin.Context{
			Request: &http.Request{
				Method: http.MethodPost,
				Header: http.Header{
					csrfHeader: []string{csrfToken},
				},
				URL: &url.URL{
					Path: "/api/v1/contracts/abcxyz",
				},
			},
		}

		csrfRepository := mockRepository.NewMockCSRFRepository(mockCtrl)
		csrfRepository.EXPECT().Get(ctx, mfidUserID, csrfToken).Return(csrfTokenStored, nil)

		v := csrfValidator{
			csrfRepository: csrfRepository,
		}

		wantErr := appErrors.NewErrorUnauthorized()
		gotErr := v.validate(ctx, entity.AuthClaims{
			MFIDUserID: mfidUserID,
			CSRFToken:  csrfToken,
		})

		assert.Equal(t, wantErr, gotErr)
	})

	t.Run("success - ignore method", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mfidUserID := uint32(1)
		csrfToken := "csrf_token"

		ctx := &gin.Context{
			Request: &http.Request{
				Method: http.MethodOptions,
			},
		}

		csrfRepository := mockRepository.NewMockCSRFRepository(mockCtrl)

		v := csrfValidator{
			csrfRepository: csrfRepository,
		}

		gotErr := v.validate(ctx, entity.AuthClaims{
			MFIDUserID: mfidUserID,
			CSRFToken:  csrfToken,
		})

		require.NoError(t, gotErr)
	})

	t.Run("success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mfidUserID := uint32(1)
		csrfToken := "csrf_token"

		ctx := &gin.Context{
			Request: &http.Request{
				Method: http.MethodPost,
				Header: http.Header{
					csrfHeader: []string{csrfToken},
				},
				URL: &url.URL{
					Path: "/api/v1/contracts/abcxyz",
				},
			},
		}

		csrfRepository := mockRepository.NewMockCSRFRepository(mockCtrl)
		csrfRepository.EXPECT().Get(ctx, mfidUserID, csrfToken).Return(csrfToken, nil)

		v := csrfValidator{
			csrfRepository: csrfRepository,
		}

		gotErr := v.validate(ctx, entity.AuthClaims{
			MFIDUserID: mfidUserID,
			CSRFToken:  csrfToken,
		})

		require.NoError(t, gotErr)
	})
}

func TestNewTermsOfUseValidator(t *testing.T) {
	tests := []struct {
		name string
		want ClaimsValidator
	}{
		{
			name: "success",
			want: &termsOfUseValidator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTermsOfUseValidator()

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_termsOfUseValidator_validate(t *testing.T) {
	type args struct {
		ctx    *gin.Context
		claims entity.AuthClaims
	}

	tests := []struct {
		name    string
		v       *termsOfUseValidator
		args    args
		wantErr error
	}{
		{
			name: "confirmed",
			v:    &termsOfUseValidator{},
			args: args{
				ctx: &gin.Context{},
				claims: entity.AuthClaims{
					AcceptedTerms: true,
				},
			},
			wantErr: nil,
		},
		{
			name: "yet confirmed",
			v:    &termsOfUseValidator{},
			args: args{
				ctx: &gin.Context{},
				claims: entity.AuthClaims{
					AcceptedTerms: false,
				},
			},
			wantErr: appErrors.NewCustomErrorForbidden(
				appErrors.CodeMFIDUserNeedToAcceptTermsOfUse,
				"mfid user needs to accept terms of use",
				nil,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &termsOfUseValidator{}
			gotErr := v.validate(tt.args.ctx, tt.args.claims)

			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
