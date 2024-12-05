package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
	"github.com/education-english-web/BE-english-web/app/errors"
	apiContext "github.com/education-english-web/BE-english-web/app/interface/api/context"
	"github.com/education-english-web/BE-english-web/app/interface/api/presenter"
	"github.com/education-english-web/BE-english-web/app/services"
	"github.com/education-english-web/BE-english-web/pkg/tracer"
)

// token keys
const (
	AccessKey      = "education_jwt_access"
	RefreshKey     = "education_jwt_refresh"
	csrfHeader     = "X-Csrf-Token"
	cookieMaxAge   = 86400
	cookiePath     = "/api"
	cookieHTTPOnly = true
)

var (
	cookieSameSite http.SameSite
	cookieSecure   bool
)

// InitCookieOptions initializes cookieOption
func InitCookieOptions(sameSite http.SameSite, secure bool) {
	cookieSameSite = sameSite
	cookieSecure = secure
}

// CookieAuthenticator is for User/Office Authentication and CSRF protection
type CookieAuthenticator struct {
	jwt        services.JWT
	validators []ClaimsValidator
}

// NewCookieAuthenticator returns CookieAuthenticator
func NewCookieAuthenticator(jwt services.JWT, validators ...ClaimsValidator) CookieAuthenticator {
	return CookieAuthenticator{
		jwt:        jwt,
		validators: validators,
	}
}

// Authenticate validates JWT and CSRF token, set userID and officeID
// this function is likely to replace Authenticate function because we will not check error in context anymore
func (m *CookieAuthenticator) Authenticate(ctx *gin.Context) {
	token, err := ctx.Cookie(AccessKey)
	if err != nil {
		presenter.RenderErrors(ctx, tracer.NoopSpan{}, errors.NewErrorUnauthorized())
		ctx.Abort()

		return
	}

	claims := entity.AuthClaims{}
	if err := m.jwt.Decrypt(token, &claims, false); err != nil {
		presenter.RenderErrors(ctx, tracer.NoopSpan{}, errors.NewErrorUnauthorized())
		ctx.Abort()

		return
	}

	// execute all validators
	for _, v := range m.validators {
		err = v.validate(ctx, claims)
		if err != nil {
			presenter.RenderErrors(ctx, tracer.NoopSpan{}, err)
			ctx.Abort()

			return
		}
	}

	//apiContext.SetMFIDUserID(ctx, claims.MFIDUserID)
	apiContext.SetUserID(ctx, claims.UserID)
	//apiContext.SetNavisOfficeID(ctx, claims.OfficeID)
	//apiContext.SetTenantUID(ctx, claims.TenantUID)

	//customCtx := context.WithValue(ctx.Request.Context(), datadog.KeyOfficeID, claims.OfficeID)
	//customCtx = context.WithValue(customCtx, datadog.KeyTenantUID, claims.TenantUID)
	//ctx.Request = ctx.Request.WithContext(customCtx)

	if claims.ProxyLoginEventID != nil {
		apiContext.SetProxyLoginEventID(ctx, *claims.ProxyLoginEventID)
	}

	ctx.Next()
}

// SetCookie is expected to create a session cookie
func setCookie(ctx *gin.Context, key, value string) {
	ctx.SetSameSite(cookieSameSite)
	ctx.SetCookie(
		key,
		value,
		cookieMaxAge,
		cookiePath,
		"",
		cookieSecure,
		cookieHTTPOnly,
	)
}

// SetCookie is expected to create a session cookie
func deleteCookie(ctx *gin.Context, key string) {
	ctx.SetCookie(
		key,
		"",
		-1,
		cookiePath,
		"",
		cookieSecure,
		cookieHTTPOnly,
	)
}

// SetAccessCookie is expected to create & refresh JWT. Ex: SessionHandler
func SetAccessCookie(ctx *gin.Context, token string) {
	setCookie(ctx, AccessKey, token)
}

func DeleteAccessCookie(ctx *gin.Context) {
	deleteCookie(ctx, AccessKey)
}

// SetRefreshCookie is expected to create & refresh JWT. Ex: SessionHandler
func SetRefreshCookie(ctx *gin.Context, token string) {
	setCookie(ctx, RefreshKey, token)
}

func DeleteRefreshCookie(ctx *gin.Context) {
	deleteCookie(ctx, RefreshKey)
}

func GetAllCookies(ctx *gin.Context) []*http.Cookie {
	return ctx.Request.Cookies()
}

func PrintAllCookies(ctx *gin.Context) {
	cookies := GetAllCookies(ctx)
	for _, cookie := range cookies {
		fmt.Printf("Cookie: %s=%s\n", cookie, cookie.Value)
	}
}
