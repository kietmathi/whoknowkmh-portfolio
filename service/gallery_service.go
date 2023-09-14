package service

import (
	"sync"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

var once sync.Once

// galleryService : This struct implements the methods defined in the domain.GalleryService
type galleryService struct {
	photoRepository domain.PhotoRepository
}

var instance *galleryService

// NewGalleryService : Create a new instance of domain.GalleryService.
// This instance receives a domain.PhotoRepository injection.
func NewGalleryService(r domain.PhotoRepository) domain.GalleryService {
	once.Do(func() {
		instance = &galleryService{
			photoRepository: r,
		}
	})
	return instance
}

// FindPhotoByID: Find the photo with specificed ID
func (s *galleryService) FindPhotoByID(id uint) (domain.Photo, error) {
	return s.photoRepository.FindByID(id)
}

// FindAllPhoto: Find all available photos and sort them in descending order
func (s *galleryService) FindAllPhoto() ([]domain.Photo, error) {
	return s.photoRepository.FindAll()
}

// FindNextAndPrevPhotoID : Find the next and previous photo IDs related to a specific ID
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
