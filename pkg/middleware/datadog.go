package middleware

import (
	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

// Constant for tracer engine
const (
	TracerEngineNop     = "nop"
	TracerEngineDatadog = "datadog"
)

// Constant for profiler engine
const (
	ProfilerEngineNop     = "nop"
	ProfilerEngineDatadog = "datadog"
)

// Nop return a middleware it does nothing
func Nop() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// DatadogTracer return a middleware for datadog tracer
// @Param serviceName the name will appear on datadog monitor
func DatadogTracer(serviceName string) gin.HandlerFunc {
	return gintrace.Middleware(serviceName, gintrace.WithIgnoreRequest(func(c *gin.Context) bool {
		return c.Request.URL.Path == "/api/healthz"
	}))
}
