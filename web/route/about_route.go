package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

// NewAboutRouter: Set up routing so that each request is directed to the 'about' controller for processing
func NewAboutRouter(gin *gin.RouterGroup) {
	abc := &controller.AboutController{
		AboutUsecase: usecase.NewAboutUsecase(),
	}

	gin.GET("/about", abc.Show)
}
