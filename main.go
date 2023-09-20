package main

import (
	"embed"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/bootstrap"
	"github.com/kietmathi/whoknowkmh-portfolio/web/middleware"
	"github.com/kietmathi/whoknowkmh-portfolio/web/route"
)

const (
	sessionsName = "mysession"
	assetsPath   = "/assets"
)

//go:embed templates/* assets
var EmbedFS embed.FS

func main() {
	// Initialize application
	app := bootstrap.App(EmbedFS)
	// environment variable
	env := app.Env
	// Instance for SQLite database
	db := app.SQLiteDB
	defer app.CloseDBConnection()

	// Initialize Gin instance
	gin := gin.Default()
	{
		// Applying middleware for the Application
		gin.Use(middleware.Sessions(sessionsName, env.SessionSecret))
		// Sử dụng middleware CSRF cho tất cả các route
		gin.Use(middleware.CSRF(env.CSRFAuthKey))
		gin.Use(middleware.CSRFToken())
		gin.Use(gzip.Gzip(gzip.DefaultCompression))
		gin.Use(cors.Default())
		gin.Use(middleware.CacheStaticFiles(2 * time.Hour))
		// Set templates
		gin.HTMLRender = app.EmbedTemplates
		// Serve static files
		assets := app.EmbedAssets
		gin.StaticFS(assetsPath, http.FS(assets))

		gin.NoRoute(middleware.NotFound())
		gin.Use(middleware.InternalServerError())

	}

	// Setup : Implement routing to direct each request to its respective controller
	route.Setup(db, env, app.Logger, gin)

	// Launch web application
	// Run the server with the specific port number
	gin.Run(env.ServerAddress)
}
