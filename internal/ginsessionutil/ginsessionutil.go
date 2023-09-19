package ginsessionutil

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Set(c context.Context, key string, value interface{}) error {
	session := sessions.Default(c.(*gin.Context))
	session.Set(key, value)
	return session.Save()
}

func Get(c context.Context, key string) (interface{}, error) {
	session := sessions.Default(c.(*gin.Context))
	return session.Get(key), nil
}

func Delete(c context.Context, key string) error {
	session := sessions.Default(c.(*gin.Context))
	session.Delete(key)
	return session.Save()
}
