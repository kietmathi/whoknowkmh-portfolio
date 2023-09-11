package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

func NewBlogRouter(gin *gin.RouterGroup) {
	bc := controller.NewBlogController()

	gin.GET("/blog", bc.Show)
}
