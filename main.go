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
	// Applying middleware for the Application
	gin.Use(gzip.Gzip(gzip.DefaultCompression))
	gin.Use(cors.Default())
	gin.Use(middleware.CacheStaticFiles(2 * time.Hour))
	// Set templates
	gin.HTMLRender = app.EmbedTemplates
	// Serve static files
	assets := app.EmbedAssets
	gin.StaticFS("/assets", http.FS(assets))

	// Setup : Implement routing to direct each request to its respective controller
	route.Setup(db, app.Logger, gin)

	// Launch web application
	// Run the server with the specific port number
	gin.Run(env.ServerAddress)
}
