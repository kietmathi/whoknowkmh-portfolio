package usecase

import (
	"context"
	"sync"
	"time"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderhelper"
)

var adminUsecaseOnce sync.Once

// galleryService : This struct implements the methods defined in the domain.GalleryService
type adminUsecase struct {
	photoRepository domain.PhotoRepository
}

var adminUsecaseInstance *adminUsecase

func NewAdminUsecase(pr domain.PhotoRepository) domain.AdminUsecase {
	adminUsecaseOnce.Do(func() {
		adminUsecaseInstance = &adminUsecase{
			photoRepository: pr,
		}
	})
	return adminUsecaseInstance
}

func (as *adminUsecase) FindAvailableDBTable() []string {
	tableName := []string{}
	tableName = append(tableName, as.photoRepository.TableName())
	return tableName
}
func (as *adminUsecase) ShowAllPhoto() ([]domain.Photo, error) {
	return as.photoRepository.FindAll()
}

func (as *adminUsecase) UpdatePhotoByID(photo domain.Photo) (domain.Photo, error) {
	photo.UpdatedAt = time.Now()
	return as.photoRepository.UpdateByID(photo)
}

func (as *adminUsecase) InsertPhoto(photo domain.Photo) (domain.Photo, error) {
	insertPhoto := domain.Photo{
		Name:        photo.Name,
		Url:         photo.Url,
		Description: photo.Description,
	}
	return as.photoRepository.InsertPhoto(insertPhoto)
}

func (as *adminUsecase) RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{}) {
	renderhelper.RenderTemplate(c, statusCode, name, cacheDuration, data)
}
