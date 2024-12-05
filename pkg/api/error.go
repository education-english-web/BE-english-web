package api

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Resp *http.Response
	Body []byte
}

func (e HTTPError) StatusCode() int {
	return e.Resp.StatusCode
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("http error: %s %s", e.Resp.Status, e.Body)
}
