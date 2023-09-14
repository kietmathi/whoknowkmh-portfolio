package renderutil

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RenderTemplte(
	c *gin.Context,
	statusCode int,
	name string,
	cacheTime time.Duration,
	data interface{}) {
	c.Header("Cache-Control", "public, max-age="+cacheTime.String())
	c.HTML(
		statusCode,
		name,
		gin.H{
			"data": data,
		})
}
