package renderutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderTemplte(
	c *gin.Context,
	statusCode int,
	name string,
	data interface{}) {
	c.HTML(
		http.StatusBadRequest,
		name,
		gin.H{
			"data": data,
		})
}
