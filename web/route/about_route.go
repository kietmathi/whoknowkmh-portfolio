package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

func NewAboutRouter(gin *gin.RouterGroup) {
	ac := controller.NewAboutController()

	gin.GET("/about", ac.Show)
}
