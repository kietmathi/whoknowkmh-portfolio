package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

// NewBlogRouter: Set up routing so that each request is directed to the 'blog' controller for processing
func NewBlogRouter(gin *gin.RouterGroup) {
	bc := &controller.BlogController{
		BlogUsecase: usecase.NewBlogUsecase(),
	}

	gin.GET("/blog", bc.Show)
}
