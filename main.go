package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/bootstrap"
	"github.com/kietmathi/whoknowkmh-portfolio/web/route"
)

//go:embed templates/* assets
var EmbedFS embed.FS

func main() {
	app := bootstrap.App(EmbedFS)

	env := app.Env

	db := app.SQLite
	defer app.CloseDBConnection()

	gin := gin.Default()

	gin.HTMLRender = app.EmbedTemplates

	assets := app.EmbedAssets
	gin.StaticFS("/assets", http.FS(assets))

	route.Setup(db, gin)

	gin.Run(env.ServerAddress)
}
