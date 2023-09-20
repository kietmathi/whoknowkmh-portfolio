package domain

import (
	"context"
	"time"
)

const (
	AboutTitle        = "about"
	AboutTemplateName = "user/about.html"
)

// AboutUsecase : represent the about's usecase
type AboutUsecase interface {
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
