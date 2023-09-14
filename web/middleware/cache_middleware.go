package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CacheStaticFiles(maxAge time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Apply the Cache-Control header to the static files
		if strings.HasPrefix(c.Request.URL.Path, "/assets/") {
			c.Header("Cache-Control", "public, max-age="+maxAge.String())
		}
		// Continue to the next middleware or handler
		c.Next()
	}
}
