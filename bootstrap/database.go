package bootstrap

import (
	"log"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

func NewSQLiteDatabase(DNS string) *gorm.DB {
	var db *gorm.DB
	db, err := gorm.Open(sqlite.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.Photo{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseSQLiteDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbSQL.Close()
}
