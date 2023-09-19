package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/ginsessionutil"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/tokenutil"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, err := c.Cookie("Authorization")
		if err != nil {
			ginsessionutil.Set(c, "error", "Not authorized")
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		authorized, err := tokenutil.IsAuthorized(authToken, secret)
		if err != nil {
			ginsessionutil.Set(c, "error", err.Error())
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		if !authorized {
			ginsessionutil.Set(c, "error", "Not authorized")
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
		if err != nil {
			ginsessionutil.Set(c, "error", err.Error())
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		c.Set("x-user-id", userID)
		c.Next()
	}
}
