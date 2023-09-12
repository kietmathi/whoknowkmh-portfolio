package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, logger *log.Logger, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewHomeRouter(db, publicRouter)
	NewAboutRouter(publicRouter)
	NewGalleryRouter(db, logger, publicRouter)
	NewBlogRouter(publicRouter)
}
