package bootstrap

import (
	"log"

	"github.com/kietmathi/whoknowkmh-portfolio/repository"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

func NewSQLiteDatabase(DNS string) *gorm.DB {
	var db *gorm.DB
	db, err := gorm.Open(sqlite.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	pr := repository.NewPhotoRepository(db)
	err = pr.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ClosSQLiteDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbSQL.Close()
}
