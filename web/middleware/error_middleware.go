package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderhelper"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		renderhelper.RenderTemplate(
			c,
			http.StatusNotFound,
			"user/not_found.html",
			0*time.Second,
			"",
		)
	}
}

func InternalServerError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				renderhelper.RenderTemplate(
					c,
					http.StatusNotFound,
					"user/internal_server_error.html",
					0*time.Second,
					"",
				)
			}
		}()
		c.Next()
	}
}
