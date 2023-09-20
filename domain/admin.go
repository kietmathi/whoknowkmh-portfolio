package domain

import (
	"context"
	"time"
)

const (
	AdminTitle = "admin"
	AdminTemplateName = "admin/admin.html"
)

// AdminUsecase : represent the admin's usecase
type AdminUsecase interface {
	FindAvailableDBTable() []string
	ShowAllPhoto() ([]Photo, error)
	UpdatePhotoByID(Photo) (Photo, error)
	InsertPhoto(Photo) (Photo, error)
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
