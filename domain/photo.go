package domain

import "time"

// Photo : Database model for photo
type Photo struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:255;not null"`
	Url         string `gorm:"size:255;not null"`
	Description string `gorm:"size:1000"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// PhotoService : represent the photo's services
type PhotoService interface {
	FindByID(uint) (Photo, error)
	FindAll() ([]Photo, error)
	GetAllID() ([]string, error)
}

// PhotoRepository : represent the Photo's repository contract
type PhotoRepository interface {
	FindByID(uint) (Photo, error)
	FindAll() ([]Photo, error)
	GetAllID() ([]string, error)
	Migrate() error
}
