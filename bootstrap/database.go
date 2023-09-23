package bootstrap

import (
	"log"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/sqlite"
)

// NewSQLiteDatabase : Create a new SQLite instance
func NewSQLiteDatabase(DNS string) sqlite.Client {
	client, err := sqlite.NewClient(DNS)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Database().AutoMigrate(&domain.Photo{})
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// CloseSQLiteDatabaseConnection: Close SQLite database connection
func CloseSQLiteDatabaseConnection(client sqlite.Client) {
	if client == nil {
		return
	}
	err := client.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}
