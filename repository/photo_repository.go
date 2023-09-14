package repository

import (
	"fmt"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"gorm.io/gorm"
)

// photoRepository: This struct implements the methods defined in the PhotoRepository.
// These methods directly manipulate the database
type photoRepository struct {
	DB *gorm.DB
}

// NewPhotoRepository : Create a new instance of PhotoRepository.
// This instance receives a database injection.
func NewPhotoRepository(db *gorm.DB) domain.PhotoRepository {
	return &photoRepository{
		DB: db,
	}
}

// FindByID: find the photo with specificed ID
func (p *photoRepository) FindByID(id uint) (domain.Photo, error) {
	var photo domain.Photo
	err := p.DB.First(&photo, fmt.Sprintf("id=%v", id)).Error
	return photo, err
}

// FindAll: Find all available photos and sort them in descending order
func (p *photoRepository) FindAll() ([]domain.Photo, error) {
	var photos []domain.Photo
	err := p.DB.Order("id DESC").Find(&photos).Error
	return photos, err
}

// GetAllID: Get all available photo IDs
func (p *photoRepository) GetAllID() ([]string, error) {
	var ids []string
	err := p.DB.Model(&domain.Photo{}).Order("id DESC").Pluck("id", &ids).Error
	return ids, err
}
