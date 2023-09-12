package service

import (
	"sync"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

var once sync.Once

type galleryService struct {
	photoRepository domain.PhotoRepository
}

var instance *galleryService

func NewGalleryService(r domain.PhotoRepository) domain.GalleryService {
	once.Do(func() {
		instance = &galleryService{
			photoRepository: r,
		}
	})
	return instance
}

func (s *galleryService) FindPhotoByID(id uint) (domain.Photo, error) {
	return s.photoRepository.FindByID(id)
}

func (s *galleryService) FindAllPhoto() ([]domain.Photo, error) {
	return s.photoRepository.FindAll()
}

func (s *galleryService) FindNextAndPrevPhotoID(targetID string) (string, string, error) {
	var prevPhotoID, nextPhotoID string = "", ""

	ids, err := s.photoRepository.GetAllID()
	if err != nil {
		return prevPhotoID, nextPhotoID, err
	}

	for i, v := range ids {
		if v == targetID {
			if i > 0 {
				prevPhotoID = ids[i-1]
			}
			if i < len(ids)-1 {
				nextPhotoID = ids[i+1]
			}
			return prevPhotoID, nextPhotoID, nil
		}
	}

	return prevPhotoID, nextPhotoID, nil
}
