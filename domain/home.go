package domain

import (
	"context"
	"time"
)

type HomeUsecase interface {
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
