package renderhelper

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

// Use the HTML method of the Context to render a template
func RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{}) {
	// Set a specific caching duration for rendered HTML content
	c.(*gin.Context).Header("Cache-Control", "public, max-age="+cacheDuration.String())
	// Set data for rendering HTML content
	c.(*gin.Context).HTML(
		// Http status code
		statusCode,
		// Template name
		name,
		// Data that the template uses for rendering
		gin.H{
			"data":           data,
			csrf.TemplateTag: csrf.TemplateField(c.(*gin.Context).Request),
		})
}
