package repository

import (
	"fmt"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"gorm.io/gorm"
)

type photoRepository struct {
	DB *gorm.DB
}

// NewPhotoRepository : get injected database
func NewPhotoRepository(db *gorm.DB) domain.PhotoRepository {
	return &photoRepository{
		DB: db,
	}
}

func (p *photoRepository) FindByID(id uint) (domain.Photo, error) {
	var photo domain.Photo
	err := p.DB.First(&photo, fmt.Sprintf("id=%v", id)).Error
	return photo, err
}

func (p *photoRepository) FindAll() ([]domain.Photo, error) {
	var photos []domain.Photo
	err := p.DB.Order("id DESC").Find(&photos).Error
	return photos, err
}

func (p *photoRepository) GetAllID() ([]string, error) {
	var ids []string
	err := p.DB.Model(&domain.Photo{}).Order("id DESC").Pluck("id", &ids).Error
	return ids, err
}

func (p *photoRepository) Migrate() error {
	return p.DB.AutoMigrate(&domain.Photo{})
}
