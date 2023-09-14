package renderutil

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Use the HTML method of the Context to render a template
func RenderTemplte(c *gin.Context, statusCode int, name string, cacheDuration time.Duration, data interface{}) {
	// Set a specific caching duration for rendered HTML content
	c.Header("Cache-Control", "public, max-age="+cacheDuration.String())
	// Set data for rendering HTML content
	c.HTML(
		// Http status code
		statusCode,
		// Template name
		name,
		// Data that the template uses for rendering
		gin.H{
			"data": data,
		})
}
