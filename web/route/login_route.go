package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/bootstrap"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

func NewLoginRouter(env *bootstrap.Env, logger *log.Logger, gin *gin.RouterGroup) {
	lgc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(),
		Env:          env,
	}

	gin.GET("/login", lgc.Login)
	gin.POST("/login", lgc.LoginPost)
}
