package usecase

import (
	"context"
	"sync"
	"time"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderhelper"
)

var aboutOnce sync.Once

type aboutUsecase struct{}

var instanceAboutUsecase *aboutUsecase

func NewAboutUsecase() domain.AboutUsecase {
	aboutOnce.Do(func() {
		instanceAboutUsecase = &aboutUsecase{}
	})
	return instanceAboutUsecase
}

func (au *aboutUsecase) RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{}) {
	renderhelper.RenderTemplate(c, statusCode, name, cacheDuration, data)
}
