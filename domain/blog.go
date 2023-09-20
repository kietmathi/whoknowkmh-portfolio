package domain

import (
	"context"
	"time"
)

const (
	BlogTitle        = "blog"
	BlogTemplateName = "user/blog.html"
)

// BlogUsecase : represent the blog's usecase
type BlogUsecase interface {
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
