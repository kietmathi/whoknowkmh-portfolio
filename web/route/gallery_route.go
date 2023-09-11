package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/repository"
	"github.com/kietmathi/whoknowkmh-portfolio/service"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
	"gorm.io/gorm"
)

func NewGalleryRouter(db *gorm.DB, gin *gin.RouterGroup) {
	pr := repository.NewPhotoRepository(db)
	ps := service.NewPhotoService(pr)
	pc := controller.NewPhotosController(ps)

	gin.GET("/gallery", pc.ShowAll)
	gin.GET("/gallery/:imgid", pc.ShowByID)
}
