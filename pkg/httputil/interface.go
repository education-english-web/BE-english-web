package httputil

import "net/http"

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type HTTPFactory interface {
	GetIP(r *http.Request) (string, error)
	GetUserAgent(r *http.Request) string
}
