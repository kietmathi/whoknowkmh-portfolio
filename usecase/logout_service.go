package usecase

import (
	"context"
	"sync"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/gincookiesessionutil"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/ginsessionutil"
)

var logoutOnce sync.Once

type logoutUsecase struct{}

var instanceLogoutUsecase *logoutUsecase

func NewLogoutUsecase() domain.LogoutUsecase {
	logoutOnce.Do(func() {
		instanceLogoutUsecase = &logoutUsecase{}
	})
	return instanceLogoutUsecase
}

func (abs *logoutUsecase) SetSession(c context.Context, key string, value interface{}) error {
	return ginsessionutil.Set(c, key, value)
}

func (abs *logoutUsecase) DeleteFromCookieSession(c context.Context, key string) error {
	return gincookiesessionutil.Delete(c, key)
}
