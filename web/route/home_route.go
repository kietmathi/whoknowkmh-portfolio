package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/repository"
	"github.com/kietmathi/whoknowkmh-portfolio/service"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
	"gorm.io/gorm"
)

func NewHomeRouter(db *gorm.DB, group *gin.RouterGroup) {
	pr := repository.NewPhotoRepository(db)
	ps := service.NewPhotoService(pr)
	pc := controller.NewHomeController(ps)

	group.GET("/", pc.Show)
}
