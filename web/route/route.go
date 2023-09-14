package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Setup : Implement routing to direct each request to its respective controller
func Setup(db *gorm.DB, logger *log.Logger, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewHomeRouter(db, publicRouter)
	NewAboutRouter(publicRouter)
	NewGalleryRouter(db, logger, publicRouter)
	NewBlogRouter(publicRouter)
}
