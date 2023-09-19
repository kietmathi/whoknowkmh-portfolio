package usecase

import (
	"context"
	"sync"
	"time"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderhelper"
)

var blogOnce sync.Once

type blogUsecase struct{}

var instanceBlogUsecase *blogUsecase

func NewBlogUsecase() domain.BlogUsecase {
	blogOnce.Do(func() {
		instanceBlogUsecase = &blogUsecase{}
	})
	return instanceBlogUsecase
}

func (bs *blogUsecase) RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{}) {
	renderhelper.RenderTemplate(c, statusCode, name, cacheDuration, data)
}
