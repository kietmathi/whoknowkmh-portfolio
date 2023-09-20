package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

func NewLogoutRouter(gin *gin.RouterGroup, logger *log.Logger) {
	loc := &controller.LogoutController{
		LogoutUsecase: usecase.NewLogoutUsecase(),
	}

	gin.GET("/logout", loc.Logout)
}
