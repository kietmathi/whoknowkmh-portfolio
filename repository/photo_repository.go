package repository

import (
	"fmt"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/sqlite"
)

// photoRepository: This struct implements the methods defined in the PhotoRepository.
// These methods directly manipulate the database
type photoRepository struct {
	db sqlite.Database
}

// NewPhotoRepository : Create a new instance of PhotoRepository.
// This instance receives a database injection.
func NewPhotoRepository(db sqlite.Database) domain.PhotoRepository {
	return &photoRepository{
		db: db,
	}
}

func (p *photoRepository) TableName() string {
	return "photo"
}

// FindByID: find the photo with specificed ID
func (p *photoRepository) FindByID(id uint) (domain.Photo, error) {
	var photo domain.Photo
	err := p.db.Model(&domain.Photo{}).First(&photo, fmt.Sprintf("id=%v", id))
	return photo, err
}

// FindAll: Find all available photos and sort them in descending order
func (p *photoRepository) FindAllAvailable() ([]domain.Photo, error) {
	var photos []domain.Photo
	err := p.db.Model(&domain.Photo{}).Where("delete_flag = ?", false).Order("id DESC").Find(&photos)
	return photos, err
}

// GetAllID: Get all available photo IDs
func (p *photoRepository) GetAllAvailableID() ([]string, error) {
	var ids []string
	err := p.db.Model(&domain.Photo{}).Select("id").Where("delete_flag = ?", false).Order("id DESC").Find(&ids)
	return ids, err
}

// FindAll: Find all available photos and sort them in descending order
func (p *photoRepository) FindAll() ([]domain.Photo, error) {
	var photos []domain.Photo
	err := p.db.Model(&domain.Photo{}).Order("id DESC").Find(&photos)
	return photos, err
}

func (p *photoRepository) UpdateByID(photo domain.Photo) (domain.Photo, error) {
	err := p.db.Model(&domain.Photo{}).Where("id = ?", photo.ID).Updates(&photo)
	return photo, err
}

func (p *photoRepository) InsertPhoto(photo domain.Photo) (domain.Photo, error) {
	err := p.db.Model(&domain.Photo{}).Create(&photo)

	return photo, err
}
