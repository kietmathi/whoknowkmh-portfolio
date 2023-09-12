package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
	"gorm.io/gorm"
)

func NewHomeRouter(db *gorm.DB, group *gin.RouterGroup) {
	pc := controller.NewHomeController()

	group.GET("/", pc.Show)
}
