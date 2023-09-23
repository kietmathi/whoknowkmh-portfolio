package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/repository"
	"github.com/kietmathi/whoknowkmh-portfolio/sqlite"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
)

// NewGalleryRouter: Set up routing so that each request is directed to the 'gallery' controller for processing
func NewGalleryRouter(db sqlite.Database, logger *log.Logger, gin *gin.RouterGroup) {
	pr := repository.NewPhotoRepository(db)
	gc := &controller.GalleryController{
		GalleryUsecase: usecase.NewGalleryUsecase(pr),
		Logger:         logger,
	}

	gin.GET("/gallery", gc.ShowAll)
	gin.GET("/gallery/:imgid", gc.ShowByID)
}
