package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/repository"
	"github.com/kietmathi/whoknowkmh-portfolio/sqlite"
	usecase "github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

func NewAdminRouter(db sqlite.Database, logger *log.Logger, gin *gin.RouterGroup) {
	pr := repository.NewPhotoRepository(db)
	ac := &controller.AdminController{
		AdminUsecase: usecase.NewAdminUsecase(pr),
		Logger:       logger,
	}

	gin.GET("/admin", ac.Show)
	gin.GET("/admin/photo", ac.ShowTablePhoto)
	gin.PUT("/admin/photo/:id", ac.UpdatePhoto)
	gin.POST("/admin/photo/:id", ac.InsertPhoto)
}
