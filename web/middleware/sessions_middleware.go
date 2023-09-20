package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/gincookiesessionutil"
)

func Sessions(name, secret string) gin.HandlerFunc {
	return sessions.Sessions(name, gincookiesessionutil.NewStore(secret))
}
