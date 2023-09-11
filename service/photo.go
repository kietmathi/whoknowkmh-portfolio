package service

import (
	"sync"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

var once sync.Once

type photoService struct {
	photoRepository domain.PhotoRepository
}

var instance *photoService

func NewPhotoService(r domain.PhotoRepository) domain.PhotoService {
	once.Do(func() {
		instance = &photoService{
			photoRepository: r,
		}
	})
	return instance
}

func (s *photoService) FindByID(id uint) (domain.Photo, error) {
	return s.photoRepository.FindByID(id)
}

func (s *photoService) FindAll() ([]domain.Photo, error) {
	return s.photoRepository.FindAll()
}

func (s *photoService) GetAllID() ([]string, error) {
	return s.photoRepository.GetAllID()
}
