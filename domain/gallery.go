package domain

import (
	"context"
	"time"
)

// PhotoService : represent the gallery's services
type GalleryUsecase interface {
	FindPhotoByID(uint) (Photo, error)
	FindAllAvailablePhoto() ([]Photo, error)
	FindNextAndPrevPhotoID(string) (string, string, error)
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
