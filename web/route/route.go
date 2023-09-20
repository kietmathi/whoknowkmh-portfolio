package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/bootstrap"
	"github.com/kietmathi/whoknowkmh-portfolio/web/middleware"
	"gorm.io/gorm"
)

// Setup : Implement routing to direct each request to its respective controller
func Setup(db *gorm.DB, env *bootstrap.Env, logger *log.Logger, gin *gin.Engine) {
	// All Public APIs
	publicRouter := gin.Group("")
	{
		NewHomeRouter(publicRouter)
		NewAboutRouter(publicRouter)
		NewGalleryRouter(db, logger, publicRouter)
		NewBlogRouter(publicRouter)
		NewLoginRouter(env, logger, publicRouter)
		NewLogoutRouter(publicRouter, logger)
	}

	protectedRouter := gin.Group("")
	{
		protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
		NewAdminRouter(db, logger, protectedRouter)
	}
}
