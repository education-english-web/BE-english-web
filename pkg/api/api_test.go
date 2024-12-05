package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/education-english-web/BE-english-web/pkg/api/mock"
)

func TestCall(t *testing.T) {
	t.Parallel()

	t.Run("make request failed", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()

		client := mock.NewMockClient(mockCtrl)

		req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:30012/internal/import", http.NoBody)
		client.EXPECT().Do(req).Return(nil, errors.New("request failed"))

		wantErr := fmt.Errorf("error while making request: %w", errors.New("request failed"))
		_, _, err := Call(ctx, client, http.MethodPost, "http://localhost:30012/internal/import", http.NoBody)
		if err == nil || err.Error() != wantErr.Error() {
			t.Errorf("Client.GetDeletedAccounts error want:%v - got:%v", wantErr, err)
		}
	})

	t.Run("response status is not successful", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()

		client := mock.NewMockClient(mockCtrl)

		req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:30012/internal/import", http.NoBody)
		client.EXPECT().Do(req).Return(&http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       io.NopCloser(strings.NewReader(`HTTP Basic: Access denied.`)),
		}, nil)

		wantErr := HTTPError{Resp: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       io.NopCloser(strings.NewReader(`HTTP Basic: Access denied.`)),
		}, Body: []byte(`HTTP Basic: Access denied.`)}
		_, _, err := Call(ctx, client, http.MethodPost, "http://localhost:30012/internal/import", http.NoBody)
		if err == nil || err.Error() != wantErr.Error() {
			t.Errorf("Client.GetDeletedAccounts error want:%v - got:%v", wantErr, err)
		}
	})

	t.Run("success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()

		client := mock.NewMockClient(mockCtrl)

		req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:30012/internal/import", http.NoBody)
		client.EXPECT().Do(req).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data": "yoo"}`)),
		}, nil)

		gotBody, gotHeader, err := Call(ctx, client, http.MethodPost, "http://localhost:30012/internal/import", http.NoBody, nil)
		if err != nil {
			t.Errorf("Client.GetDeletedAccounts error want:nil - got:%v", err)

			return
		}
		if gotHeader != nil {
			t.Errorf("Client.GetDeletedAccounts header want:nil - got:%v", gotHeader)

			return
		}
		if string(gotBody) != `{"data": "yoo"}` {
			t.Errorf("Client.GetDeletedAccounts body want:%s - got:%s", string(gotBody), `{"data": "yoo"}`)
		}
	})

	t.Run("success with request option", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		ctx := context.Background()

		client := mock.NewMockClient(mockCtrl)

		req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:30012/internal/import", http.NoBody)
		req.SetBasicAuth("username", "password")
		client.EXPECT().Do(req).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data": "yoo"}`)),
		}, nil)

		gotBody, gotHeader, err := Call(ctx, client, http.MethodPost, "http://localhost:30012/internal/import", http.NoBody, WithBasicAuthentication("username", "password"))
		if err != nil {
			t.Errorf("Client.GetDeletedAccounts error want:nil - got:%v", err)

			return
		}
		if gotHeader != nil {
			t.Errorf("Client.GetDeletedAccounts header want:nil - got:%v", gotHeader)

			return
		}
		if string(gotBody) != `{"data": "yoo"}` {
			t.Errorf("Client.GetDeletedAccounts body want:%s - got:%s", string(gotBody), `{"data": "yoo"}`)
		}
	})
}

func Test_isSuccess(t *testing.T) {
	type args struct {
		code int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success request",
			args: args{
				code: 201,
			},
			want: true,
		},
		{
			name: "fail request",
			args: args{
				code: 409,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSuccess(tt.args.code); got != tt.want {
				t.Errorf("isSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithBasicAuthentication(t *testing.T) {
	username := "username"
	password := "password"

	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	reqOption := WithBasicAuthentication(username, password)

	reqOption(req)

	gotUsername, gotPassword, ok := req.BasicAuth()

	assert.True(t, ok)
	assert.Equal(t, username, gotUsername)
	assert.Equal(t, password, gotPassword)
}

func TestAuthorization(t *testing.T) {
	token := "token"

	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	reqOption := WithAuthorization(token)
	reqOption(req)

	gotToken := req.Header.Get("Authorization")
	assert.Equal(t, token, gotToken)
}

func TestWithContentType(t *testing.T) {
	contentType := "application/json"

	req := httptest.NewRequest(http.MethodPost, "http://localhost", nil)
	reqOption := WithContentType("application/json")
	reqOption(req)

	gotContentType := req.Header.Get("Content-Type")
	assert.Equal(t, contentType, gotContentType)
}

func TestWithHeader(t *testing.T) {
	headerValue := "1234"

	req := httptest.NewRequest(http.MethodPost, "http://localhost", nil)
	reqOption := WithHeader("x-tenant-uid", headerValue)
	reqOption(req)

	gotHeaderValue := req.Header.Get("x-tenant-uid")
	assert.Equal(t, headerValue, gotHeaderValue)
}

func TestWithQueryString(t *testing.T) {
	now := time.Now()

	qs := map[string][]string{
		"name": {"a_name_here"},
		"date": {now.String()},
	}

	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	reqOption := WithQueryString(qs)

	reqOption(req)

	want := fmt.Sprintf(
		"http://localhost?%s&%s",
		fmt.Sprintf("date=%s", url.QueryEscape(now.String())),
		"name=a_name_here",
	)

	// you may notice in the `qs`, we have `name` before `date`,
	// but the string that we want is having `date` is before `name`
	// it's because query.Encode() will sort the keys by name

	got := req.URL.String()

	if diff := cmp.Diff(got, want); diff != "" {
		t.Error(diff)
	}
}
