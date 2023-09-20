package domain

import (
	"context"
)


// LogoutUsecase : represent the logout's usecase
type LogoutUsecase interface {
	SetSession(c context.Context, key string, value interface{}) error
	DeleteFromCookieSession(c context.Context, key string) error
}
