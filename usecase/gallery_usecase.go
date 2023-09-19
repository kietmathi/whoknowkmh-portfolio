package usecase

import (
	"context"
	"sync"
	"time"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderhelper"
)

var once sync.Once

// galleryService : This struct implements the methods defined in the domain.GalleryService
type galleryUsecase struct {
	photoRepository domain.PhotoRepository
}

var instance *galleryUsecase

// NewGalleryService : Create a new instance of domain.GalleryService.
// This instance receives a domain.PhotoRepository injection.
func NewGalleryUsecase(r domain.PhotoRepository) domain.GalleryUsecase {
	once.Do(func() {
		instance = &galleryUsecase{
			photoRepository: r,
		}
	})
	return instance
}

// FindPhotoByID: Find the photo with specificed ID
func (gs *galleryUsecase) FindPhotoByID(id uint) (domain.Photo, error) {
	return gs.photoRepository.FindByID(id)
}

// FindAllPhoto: Find all available photos and sort them in descending order
func (gs *galleryUsecase) FindAllAvailablePhoto() ([]domain.Photo, error) {
	return gs.photoRepository.FindAllAvailable()
}

// FindNextAndPrevPhotoID : Find the next and previous photo IDs related to a specific ID
func (gs *galleryUsecase) FindNextAndPrevPhotoID(targetID string) (string, string, error) {
	var prevPhotoID, nextPhotoID string = "", ""

	ids, err := gs.photoRepository.GetAllAvailableID()
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

func (gs *galleryUsecase) RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{}) {
	renderhelper.RenderTemplate(c, statusCode, name, cacheDuration, data)
}
