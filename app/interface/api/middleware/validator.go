package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/domain/repository"
	"github.com/education-english-web/BE-english-web/app/errors"
)

var (
	allowMethods = []string{http.MethodGet, http.MethodHead, http.MethodOptions}
	allowPaths   = []string{
		"/api/v1",
	}
)

// ClaimsValidator contains methods for verifying claims
type ClaimsValidator interface {
	validate(ctx *gin.Context, claim entity.AuthClaims) error
}

type userValidator struct{}

// NewUserValidator return a claims validator for checking whether claims contains user id information
func NewUserValidator() ClaimsValidator {
	return &userValidator{}
}

func (v *userValidator) validate(_ *gin.Context, claim entity.AuthClaims) error {
	if claim.UserID == uuid.Nil {
		return errors.NewErrorUnauthorized()
	}

	return nil
}

type rolesValidator struct{}

// NewRolesValidator return a claims validator for checking whether claims contains roles information
func NewRolesValidator() ClaimsValidator {
	return &rolesValidator{}
}

func (v *rolesValidator) validate(_ *gin.Context, claims entity.AuthClaims) error {
	if claims.Roles <= 0 {
		return errors.NewErrorForbidden()
	}

	return nil
}

type csrfValidator struct {
	csrfRepository repository.CSRFRepository
}

// NewCSRFValidator return a claims validator for checking whether claims contain the correct csrf token
func NewCSRFValidator(csrfRepository repository.CSRFRepository) ClaimsValidator {
	return &csrfValidator{
		csrfRepository: csrfRepository,
	}
}

func (v *csrfValidator) validate(ctx *gin.Context, claims entity.AuthClaims) error {
	if !slices.Contains(allowMethods, ctx.Request.Method) &&
		!(ctx.Request.Method == http.MethodPost && slices.Contains(allowPaths, ctx.Request.URL.Path)) {
		csrfTokenStored, err := v.csrfRepository.Get(ctx, claims.UserID.String(), claims.CSRFToken)
		if err != nil {
			return errors.NewErrorUnauthorized()
		}

		csrfToken := ctx.Request.Header.Get(csrfHeader)
		if csrfTokenStored != csrfToken {
			return errors.NewErrorUnauthorized()
		}
	}

	return nil
}

type termsOfUseValidator struct{}

// NewTermsOfUseValidator return a claims validator for term of use agreement
func NewTermsOfUseValidator() ClaimsValidator {
	return &termsOfUseValidator{}
}

func (v *termsOfUseValidator) validate(_ *gin.Context, claims entity.AuthClaims) error {
	if !claims.AcceptedTerms {
		return errors.NewCustomErrorForbidden(
			errors.CodeMFIDUserNeedToAcceptTermsOfUse,
			"user needs to accept terms of use",
			nil,
		)
	}

	return nil
}
