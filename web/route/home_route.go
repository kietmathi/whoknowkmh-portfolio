package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
	"gorm.io/gorm"
)

// NewHomeRouter: Set up routing so that each request is directed to the 'home' controller for processing
func NewHomeRouter(db *gorm.DB, group *gin.RouterGroup) {
	pc := controller.NewHomeController()

	group.GET("/", pc.Show)
}
