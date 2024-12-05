package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery
// a custom recovery for Gin requests
//
// We decide to not use gin.Recovery() because it may log request headers into stderr under somce circumstances,
// which is what we don't expect.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Error(err.(error)) //nolint:errcheck,forcetypeassert

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
