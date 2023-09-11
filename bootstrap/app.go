package bootstrap

import (
	"embed"
	"io/fs"

	"github.com/gin-contrib/multitemplate"
	"gorm.io/gorm"
)

type Application struct {
	Env            *Env
	SQLite         *gorm.DB
	EmbedTemplates multitemplate.Renderer
	EmbedAssets    fs.FS
}

func App(embedFS embed.FS) Application {
	app := &Application{}
	app.Env = NewEnv(embedFS)
	app.SQLite = NewSQLiteDatabase(app.Env.DatabasePath)
	app.EmbedTemplates = NewEmbedTemplates(embedFS)
	app.EmbedAssets = NewEmbedAssets(embedFS)

	return *app
}

func (app *Application) CloseDBConnection() {
	ClosSQLiteDatabaseConnection(app.SQLite)
}
