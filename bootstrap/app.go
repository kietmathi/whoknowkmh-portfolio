package bootstrap

import (
	"embed"
	"io/fs"
	"log"

	"github.com/gin-contrib/multitemplate"
	"gorm.io/gorm"
)

type Application struct {
	Env            *Env
	SQLiteDB       *gorm.DB
	EmbedTemplates multitemplate.Renderer
	EmbedAssets    fs.FS
	Logger         *log.Logger
}

func App(embedFS embed.FS) Application {
	app := &Application{}
	app.Env = NewEnv(embedFS)
	app.SQLiteDB = NewSQLiteDatabase(app.Env.DatabasePath)
	app.EmbedTemplates = NewEmbedTemplates(embedFS)
	app.EmbedAssets = NewEmbedAssets(embedFS)
	app.Logger = log.Default()

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseSQLiteDatabaseConnection(app.SQLiteDB)
}
