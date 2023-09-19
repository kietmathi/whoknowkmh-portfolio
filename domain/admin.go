package domain

import (
	"context"
	"time"
)

// AdminService : represent the admin's services
type AdminUsecase interface {
	FindAvailableDBTable() []string
	ShowAllPhoto() ([]Photo, error)
	UpdatePhotoByID(Photo) (Photo, error)
	InsertPhoto(Photo) (Photo, error)
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
