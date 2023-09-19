package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

// NewHomeRouter: Set up routing so that each request is directed to the 'home' controller for processing
func NewHomeRouter(group *gin.RouterGroup) {
	hc := &controller.HomeController{
		HomeUsecase: usecase.NewHomeUsecase(),
	}

	group.GET("/", hc.Show)
}
