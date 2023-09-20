package domain

import (
	"context"
	"time"
)

const (
	GalleryTitle              = "blog"
	GalleryAllTemplateName    = "user/gallery_all.html"
	GallerySingleTemplateName = "user/gallery_single.html"
)

// GalleryUsecase : represent the gallery's usecase
type GalleryUsecase interface {
	FindPhotoByID(uint) (Photo, error)
	FindAllAvailablePhoto() ([]Photo, error)
	FindNextAndPrevPhotoID(string) (string, string, error)
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
