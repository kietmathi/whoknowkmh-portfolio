package domain

import (
	"context"
	"time"
)

type LoginUser struct {
	UserID   string `json:"userid" form:"userid" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginUsecase interface {
	SetSession(c context.Context, key string, value interface{}) error
	GetSession(c context.Context, key string) (interface{}, error)
	DeleteSession(c context.Context, key string) error
	CreateAccessToken(user *LoginUser, secret string, expiry int) (string, error)
	SetCookieSession(c context.Context, key string, value string, maxAge int) error
	RenderTemplate(c context.Context, statusCode int, name string, cacheDuration time.Duration, data interface{})
}
