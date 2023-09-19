package usecase

import (
	"context"
	"sync"
	"time"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/gincookiesessionutil"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/ginsessionutil"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderhelper"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/tokenutil"
)

var loginOnce sync.Once

type loginUsecase struct{}

var instanceloginUsecase *loginUsecase

func NewLoginUsecase() domain.LoginUsecase {
	loginOnce.Do(func() {
		instanceloginUsecase = &loginUsecase{}
	})
	return instanceloginUsecase
}

func (lgs *loginUsecase) SetSession(c context.Context, key string, value interface{}) error {
	return ginsessionutil.Set(c, key, value)
}

func (lgs *loginUsecase) GetSession(c context.Context, key string) (interface{}, error) {
	return ginsessionutil.Get(c, key)
}

func (lgs *loginUsecase) DeleteSession(c context.Context, key string) error {
	return ginsessionutil.Delete(c, key)
}

func (lgs *loginUsecase) SetCookieSession(c context.Context, key string, value string, maxAge int) error {
	return gincookiesessionutil.Set(c, key, value, maxAge)
}

func (lgs *loginUsecase) CreateAccessToken(user *domain.LoginUser, secret string, expiry int) (string, error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lgs *loginUsecase) RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{}) {
	renderhelper.RenderTemplate(c, statusCode, name, cacheDuration, data)
}
