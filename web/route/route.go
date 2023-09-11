package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewHomeRouter(db, publicRouter)
	NewAboutRouter(publicRouter)
	NewGalleryRouter(db, publicRouter)
	NewBlogRouter(publicRouter)
}
