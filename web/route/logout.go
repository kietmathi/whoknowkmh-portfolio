package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

func NewLogoutRouter(gin *gin.RouterGroup) {
	loc := &controller.LogoutController{
		LogoutUsecase: usecase.NewLogoutUsecase(),
	}

	gin.GET("/logout", loc.Logout)
}
