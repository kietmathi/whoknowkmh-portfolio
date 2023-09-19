package usecase

import (
	"context"
	"sync"
	"time"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderhelper"
)

var homeOnce sync.Once

type homeUsecase struct{}

var instanceHomeUsecase *homeUsecase

func NewHomeUsecase() domain.HomeUsecase {
	homeOnce.Do(func() {
		instanceHomeUsecase = &homeUsecase{}
	})
	return instanceHomeUsecase
}

func (hu *homeUsecase) RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{}) {
	renderhelper.RenderTemplate(c, statusCode, name, cacheDuration, data)
}
