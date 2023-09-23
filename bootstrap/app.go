package bootstrap

import (
	"embed"
	"io/fs"
	"log"

	"github.com/gin-contrib/multitemplate"
	"github.com/kietmathi/whoknowkmh-portfolio/sqlite"
)

type Application struct {
	Env            *Env          // Environment variable
	SQLiteDB       sqlite.Client // Instance for SQLite database
	EmbedTemplates multitemplate.Renderer
	EmbedAssets    fs.FS
	Logger         *log.Logger
}

// App : Initialize application
func App(embedFS embed.FS) Application {
	app := &Application{}
	// Set data for the application
	app.Env = NewEnv(embedFS)
	app.SQLiteDB = NewSQLiteDatabase(app.Env.DatabasePath)
	app.EmbedTemplates = NewEmbedTemplates(embedFS)
	app.EmbedAssets = NewEmbedAssets(embedFS)
	app.Logger = log.Default()

	return *app
}

// CloseDBConnection : close database connection
func (app *Application) CloseDBConnection() {
	CloseSQLiteDatabaseConnection(app.SQLiteDB)
}
