package domain

import (
	"context"
)

type LogoutUsecase interface {
	SetSession(c context.Context, key string, value interface{}) error
	DeleteFromCookieSession(c context.Context, key string) error
}
