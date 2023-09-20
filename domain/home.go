package domain

import (
	"context"
	"time"
)

const (
	HomeTitle        = "home"
	HomeTemplateName = "user/home.html"
)

// HomeUsecase : represent the home's usecase
type HomeUsecase interface {
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
