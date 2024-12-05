package api

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestOption func(req *http.Request)

func WithBasicAuthentication(username, password string) RequestOption {
	return func(req *http.Request) {
		req.SetBasicAuth(username, password)
	}
}

func WithAuthorization(token string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set("Authorization", token)
	}
}

func WithContentType(contentType string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set("Content-Type", contentType)
	}
}

func WithHeader(key, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

func WithQueryString(qs map[string][]string) RequestOption {
	return func(req *http.Request) {
		query := req.URL.Query()

		for key, values := range qs {
			for _, v := range values {
				query.Add(key, v)
			}
		}

		req.URL.RawQuery = query.Encode()
	}
}

func Call(
	ctx context.Context,
	client Client,
	method,
	url string,
	body io.Reader,
	reqOptions ...RequestOption,
) ([]byte, map[string][]string, error) {
	// initialize HTTP request
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, nil, fmt.Errorf("error while initializing http request: %w", err)
	}

	// apply request options
	for _, opt := range reqOptions {
		if !reflect.ValueOf(opt).IsNil() {
			opt(req)
		}
	}

	// make request
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error while making request: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	// read response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error while reading response body: %w", err)
	}

	// check http status code
	if code := resp.StatusCode; !isSuccess(code) {
		return nil, nil, HTTPError{Resp: resp, Body: responseBody}
	}

	return responseBody, resp.Header, nil
}

func isSuccess(code int) bool { return code >= 200 && code < 300 }

func NewClient(timeout time.Duration) Client {
	return &http.Client{
		Timeout: timeout,
	}
}
