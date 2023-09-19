package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

type LogoutController struct {
	LogoutUsecase domain.LogoutUsecase
}

func (loc *LogoutController) Logout(c *gin.Context) {
	err := loc.LogoutUsecase.DeleteFromCookieSession(c, "Authorization")
	if err != nil {
		loc.LogoutUsecase.SetSession(c, "error", err.Error())
	}
	c.Redirect(http.StatusSeeOther, "/login")
}
