package domain

import "time"

// Photo : Database model for photo table
type Photo struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Name        string    `gorm:"size:255;not null" json:"name" form:"name" binding:"required"`
	Url         string    `gorm:"size:255;not null" json:"url" form:"url" binding:"required"`
	Description string    `gorm:"size:1000" json:"description" form:"description"`
	DeleteFlag  *bool     `gorm:"default:true" json:"deleteFlag" form:"deleteFlag"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updatedAt"`
}

// PhotoRepository : represent the Photo's repository contract
type PhotoRepository interface {
	TableName() string
	FindByID(uint) (Photo, error)
	FindAllAvailable() ([]Photo, error)
	GetAllAvailableID() ([]string, error)
	FindAll() ([]Photo, error)
	UpdateByID(Photo) (Photo, error)
	InsertPhoto(Photo) (Photo, error)
}
