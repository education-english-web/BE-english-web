package oidcclient

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"

	"github.com/education-english-web/BE-english-web/app/domain/service"
	"github.com/education-english-web/BE-english-web/app/errors"
)

const (
	idTokenAttr = "id_token"
)

// oidcService interacts with identity provider using OAuth2
type oidcService struct {
	oauth2Cnf       *oauth2.Config
	verifier        *oidc.IDTokenVerifier
	issuer          string
	allowDomainsMap map[string]bool
	jwtConfig       *jwt.Config
}

func NewOIDCService(cnf *oauth2.Config, verifier *oidc.IDTokenVerifier, issuer string, allowDomains []string, jwtConfig *jwt.Config) service.OIDCService {
	return &oidcService{
		oauth2Cnf: cnf,
		verifier:  verifier,
		issuer:    issuer,
		allowDomainsMap: func() map[string]bool {
			m := make(map[string]bool)
			for _, v := range allowDomains {
				if v != "" {
					m[v] = true
				}
			}

			return m
		}(),
		jwtConfig: jwtConfig,
	}
}

// GetToken gets accessToken following OAuth2 authorization code flow
func (r *oidcService) GetToken(code, nonce, redirectURI string) (oauth2.Token, error) {
	ctx := context.Background()

	r.oauth2Cnf.RedirectURL = redirectURI

	oauth2Token, err := r.oauth2Cnf.Exchange(ctx, code)
	if err != nil {
		return oauth2.Token{}, err
	}

	rawIDToken, ok := oauth2Token.Extra(idTokenAttr).(string)
	if !ok {
		return oauth2.Token{}, fmt.Errorf("wrong id_token")
	}

	idToken, err := r.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return oauth2.Token{}, err
	}

	if idToken.Nonce != nonce {
		return oauth2.Token{}, fmt.Errorf("wrong nonce")
	}

	return *oauth2Token, nil
}

// GetClaims gets claims of OAuth2 token
func (r *oidcService) GetClaims(token oauth2.Token, claims interface{}) error {
	ctx := context.Background()

	rawIDToken, ok := token.Extra(idTokenAttr).(string)
	if !ok {
		return fmt.Errorf("wrong id_token")
	}

	idToken, err := r.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

	return idToken.Claims(&claims)
}

// RefreshToken gets accessToken following OAuth2 authorization code flow
func (r *oidcService) RefreshToken(oauth2Token oauth2.Token) (oauth2.Token, error) {
	ctx := context.Background()
	oauth2TokenSource := r.oauth2Cnf.TokenSource(ctx, &oauth2Token)

	oauth2NewToken, err := oauth2TokenSource.Token()
	if err != nil {
		return oauth2.Token{}, err
	}

	return *oauth2NewToken, nil
}

func (r *oidcService) GetTokenFromJwtBearer(username string) (oauth2.Token, error) {
	if r.jwtConfig == nil {
		return oauth2.Token{}, fmt.Errorf("oidc not setup for jwt token flow")
	}

	r.jwtConfig.Subject = username

	oauth2Token, err := r.jwtConfig.TokenSource(context.Background()).Token()
	if err != nil {
		return oauth2.Token{}, extractOauth2Error(err)
	}

	return *oauth2Token, nil
}

// GetUserInfo gets userInfo following OIDC flow
func (r *oidcService) GetUserInfo(token oauth2.Token) (service.OAuth2UserInfo, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, r.issuer)
	if err != nil {
		return service.OAuth2UserInfo{}, err
	}

	userInfo, err := provider.UserInfo(ctx, r.oauth2Cnf.TokenSource(ctx, &token))
	if err != nil {
		return service.OAuth2UserInfo{}, err
	}

	body := make(map[string]interface{})
	if err := userInfo.Claims(&body); err != nil {
		return service.OAuth2UserInfo{}, err
	}

	// applied for public identity providers
	if !r.isAllowed(userInfo.Email) {
		return service.OAuth2UserInfo{}, errors.ErrForbidden
	}

	return service.OAuth2UserInfo{
		Subject:       userInfo.Subject,
		Profile:       userInfo.Email,
		Email:         userInfo.Email,
		EmailVerified: userInfo.EmailVerified,
		Claims:        body,
	}, nil
}

func (r *oidcService) isAllowed(email string) bool {
	if len(r.allowDomainsMap) == 0 {
		return true
	}

	mailParts := strings.Split(email, "@")
	if len(mailParts) <= 1 {
		return false
	}

	return r.allowDomainsMap[mailParts[1]]
}

func extractOauth2Error(err error) error {
	e, ok := err.(*oauth2.RetrieveError)
	if ok {
		eResp := make(map[string]string)
		_ = json.Unmarshal(e.Body, &eResp)
		errStr := eResp["error"]
		errDesc := eResp["error_description"]

		if errStr == "invalid_grant" && errDesc == "user hasn't approved this consumer" {
			return errors.ErrSalesforceAppNotInstalled
		}
	}

	return e
}
