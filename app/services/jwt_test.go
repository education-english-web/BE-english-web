package services

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/google/go-cmp/cmp"
)

var (
	_secret                  = "secret_key"
	_anotherSecret           = "another_secret_key"
	_twelveHoursLater        = time.Now().Add(time.Hour * 12)
	_userID           uint32 = 1

	_OKApprovalClaims = _approvalClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: _twelveHoursLater.Unix(),
			Id:        "",
			IssuedAt:  0,
			Issuer:    "",
			NotBefore: 0,
			Subject:   "",
		},
		UserID:  _userID,
		TokenID: "01BJMVNPBBZC3E36FJTGVF0C4S",
	}
)

type _approvalClaims struct {
	jwt.StandardClaims
	UserID  uint32 `json:"user_id,omitempty"`
	TokenID string `json:"token_id,omitempty"`
}

func TestNew(t *testing.T) {
	type args struct {
		jwtSecret string
	}

	tests := []struct {
		name string
		args args
		want JWT
	}{
		{
			name: "New JWT",
			args: args{
				jwtSecret: _secret,
			},
			want: &jwtGo{
				jwtSecret: []byte(_secret),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJWT(tt.args.jwtSecret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtGo_Encrypt(t *testing.T) {
	t.Run("encrypt OK", func(t *testing.T) {
		jGo := jwtGo{
			jwtSecret: []byte(_secret),
		}
		_, err := jGo.Encrypt(_OKApprovalClaims)
		if err != nil {
			t.Errorf("did not expect error but still get %v", err)
		}
	})
}

func Test_jwtGo_Decrypt(t *testing.T) {
	t.Run("decrypt OK", func(t *testing.T) {
		jGo := jwtGo{
			jwtSecret: []byte(_secret),
		}
		token, err := jGo.Encrypt(_OKApprovalClaims)
		if err != nil {
			t.Errorf("did not expect error but still get %v", err)

			return
		}
		var parsedClaims _approvalClaims

		if err := jGo.Decrypt(token, &parsedClaims, false); err != nil {
			t.Errorf("did not expect error but still get %v", err)

			return
		}
		if diff := cmp.Diff(parsedClaims, _OKApprovalClaims); diff != "" {
			t.Error(diff)

			return
		}
	})
	t.Run("decrypt failed, not the secret key that signed", func(t *testing.T) {
		jGo1 := jwtGo{
			jwtSecret: []byte(_secret),
		}
		token, err := jGo1.Encrypt(_approvalClaims{})
		if err != nil {
			t.Errorf("did not expect error but still get %v", err)

			return
		}

		jGo2 := jwtGo{
			jwtSecret: []byte(_anotherSecret),
		}
		var parsedClaims _approvalClaims

		err = jGo2.Decrypt(token, &parsedClaims, false)
		expectedError := errors.New("signature is invalid")
		if diff := cmp.Diff(err.Error(), expectedError.Error()); diff != "" {
			t.Error(diff)

			return
		}
	})
	t.Run("decrypt failed, the claims is expired", func(t *testing.T) {
		jGo := jwtGo{
			jwtSecret: []byte(_secret),
		}
		expiredToken, err := jGo.Encrypt(_approvalClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(-24 * time.Hour).Unix(),
			},
		})
		if err != nil {
			t.Errorf("did not expect error but still get %v", err)

			return
		}

		var parsedClaims _approvalClaims
		err = jGo.Decrypt(expiredToken, &parsedClaims, false)
		if err == nil {
			t.Error("expect error but get nil error")

			return
		}
		if !strings.HasPrefix(err.Error(), "token is expired by") {
			t.Errorf("expect error %v not match", err)
		}
	})
}
