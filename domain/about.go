package domain

import (
	"context"
	"time"
)

type AboutUsecase interface {
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
