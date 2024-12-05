package middleware

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"

	"github.com/education-english-web/BE-english-web/pkg/tracer/datadog"
)

func TestLogFormatterJSON(t *testing.T) {
	type args struct {
		params gin.LogFormatterParams
	}

	req := &http.Request{}

	span, spanCtx := datadog.StartSpanFromCtx(req.Context())
	defer span.Finish()

	req = req.WithContext(spanCtx)

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - latency is smaller than a minute",
			args: args{
				params: gin.LogFormatterParams{
					TimeStamp:    time.Date(2023, time.June, 30, 1, 2, 3, 4, time.Local),
					StatusCode:   http.StatusOK,
					Latency:      time.Minute - 1,
					ClientIP:     "127.0.0.1",
					Method:       http.MethodPost,
					Path:         "/api/v1/me",
					ErrorMessage: "",
					Request:      req,
				},
			},
			want: fmt.Sprintln(`{"msg":"finish route","request_meta":{"client_ip":"127.0.0.1","error_message":"","latency":"59.999999999s","method":"POST","path":"/api/v1/me","response_time":"2023/06/30 - 01:02:03","span_id":0,"status_code":200,"trace_id":0}}`),
		},
		{
			name: "success - latency is greater than a minute",
			args: args{
				params: gin.LogFormatterParams{
					TimeStamp:    time.Date(2023, time.June, 30, 1, 2, 3, 4, time.Local),
					StatusCode:   http.StatusOK,
					Latency:      time.Minute + 1,
					ClientIP:     "127.0.0.1",
					Method:       http.MethodPost,
					Path:         "/api/v1/me",
					ErrorMessage: "",
					Request:      req,
				},
			},
			want: fmt.Sprintln(`{"msg":"finish route","request_meta":{"client_ip":"127.0.0.1","error_message":"","latency":"1m0s","method":"POST","path":"/api/v1/me","response_time":"2023/06/30 - 01:02:03","span_id":0,"status_code":200,"trace_id":0}}`),
		},
		{
			name: "success - span context not exist",
			args: args{
				params: gin.LogFormatterParams{
					TimeStamp:    time.Date(2023, time.June, 30, 1, 2, 3, 4, time.Local),
					StatusCode:   http.StatusOK,
					Latency:      time.Minute - 1,
					ClientIP:     "127.0.0.1",
					Method:       http.MethodPost,
					Path:         "/api/v1/me",
					ErrorMessage: "",
					Request:      &http.Request{},
				},
			},
			want: fmt.Sprintln(`{"msg":"finish route","request_meta":{"client_ip":"127.0.0.1","error_message":"","latency":"59.999999999s","method":"POST","path":"/api/v1/me","response_time":"2023/06/30 - 01:02:03","status_code":200}}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LogFormatterJSON(tt.args.params)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
