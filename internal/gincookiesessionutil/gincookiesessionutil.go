package gincookiesessionutil

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewStore(secret string) cookie.Store {
	return cookie.NewStore([]byte(secret))
}

func Set(c context.Context, key string, value string, maxAge int) error {
	c.(*gin.Context).SetSameSite(http.SameSiteLaxMode)
	c.(*gin.Context).SetCookie(key, value, maxAge, "", "", false, true)
	return nil
}

func Get(c context.Context, key string) (interface{}, error) {
	value, err := c.(*gin.Context).Cookie(key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func Delete(c context.Context, key string) error {
	c.(*gin.Context).SetSameSite(http.SameSiteLaxMode)
	c.(*gin.Context).SetCookie(key, "", -1, "", "", false, true)
	return nil
}
