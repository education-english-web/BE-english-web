package api

import (
	"net/http"
	"testing"
)

func TestHTTPError_StatusCode(t *testing.T) {
	type fields struct {
		Resp *http.Response
		Body []byte
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "success",
			fields: fields{
				Resp: &http.Response{
					StatusCode: http.StatusNotFound,
					Status:     "404 not found",
				},
				Body: []byte(`body`),
			},
			want: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := HTTPError{
				Resp: tt.fields.Resp,
				Body: tt.fields.Body,
			}
			if got := e.StatusCode(); got != tt.want {
				t.Errorf("HTTPError.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPError_Error(t *testing.T) {
	type fields struct {
		Resp *http.Response
		Body []byte
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				Resp: &http.Response{
					Status: "404 not found",
				},
				Body: []byte(`body`),
			},
			want: "http error: 404 not found body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := HTTPError{
				Resp: tt.fields.Resp,
				Body: tt.fields.Body,
			}

			if got := e.Error(); got != tt.want {
				t.Errorf("HTTPError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
