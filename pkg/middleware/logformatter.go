package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func LogFormatterJSON(params gin.LogFormatterParams) string {
	latency := params.Latency
	if latency > time.Minute {
		latency = latency.Truncate(time.Second)
	}

	requestMeta := map[string]interface{}{
		"response_time": params.TimeStamp.Format("2006/01/02 - 15:04:05"),
		"status_code":   params.StatusCode,
		"latency":       latency.String(),
		"client_ip":     params.ClientIP,
		"path":          params.Path,
		"method":        params.Method,
		"error_message": params.ErrorMessage,
	}

	span, isExist := tracer.SpanFromContext(params.Request.Context())
	if isExist {
		requestMeta["trace_id"] = span.Context().TraceID()
		requestMeta["span_id"] = span.Context().SpanID()
	}

	raw := map[string]interface{}{
		"request_meta": requestMeta,
		"msg":          "finish route",
	}

	//nolint:errchkjson
	bytes, _ := json.Marshal(&raw)

	return fmt.Sprintf("%s\n", string(bytes))
}
