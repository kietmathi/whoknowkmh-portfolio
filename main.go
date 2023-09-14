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
	app := bootstrap.App(EmbedFS)

	env := app.Env

	db := app.SQLiteDB
	defer app.CloseDBConnection()

	gin := gin.Default()
	gin.Use(gzip.Gzip(gzip.DefaultCompression))
	gin.Use(cors.Default())
	gin.Use(middleware.CacheStaticFiles(2 * time.Hour))
	gin.HTMLRender = app.EmbedTemplates
	assets := app.EmbedAssets
	gin.StaticFS("/assets", http.FS(assets))

	route.Setup(db, app.Logger, gin)

	gin.Run(env.ServerAddress)
}
