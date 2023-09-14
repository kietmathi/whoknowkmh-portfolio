package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

// NewAboutRouter: Set up routing so that each request is directed to the 'about' controller for processing
func NewAboutRouter(gin *gin.RouterGroup) {
	ac := controller.NewAboutController()

	gin.GET("/about", ac.Show)
}
