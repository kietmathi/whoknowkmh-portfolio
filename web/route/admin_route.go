package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/repository"
	usecase "github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
	"gorm.io/gorm"
)

func NewAdminRouter(db *gorm.DB, gin *gin.RouterGroup) {
	pr := repository.NewPhotoRepository(db)
	ac := &controller.AdminController{
		AdminUsecase: usecase.NewAdminUsecase(pr),
	}

	gin.GET("/admin", ac.Show)
	gin.GET("/admin/:tablename", ac.ShowTableDetail)
	gin.PUT("/admin/photo/:id", ac.UpdatePhoto)
	gin.POST("/admin/photo/:id", ac.InsertPhoto)
}
