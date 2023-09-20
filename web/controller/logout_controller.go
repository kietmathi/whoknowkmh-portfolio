package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

type LogoutController struct {
	LogoutUsecase domain.LogoutUsecase
	Logger        domain.Logger
}

func (loc *LogoutController) Logout(c *gin.Context) {
	err := loc.LogoutUsecase.DeleteFromCookieSession(c, "Authorization")
	if err != nil {
		loc.Logger.Printf("%+v\n", err)
		loc.LogoutUsecase.SetSession(c, "error", err.Error())
	}
	c.Redirect(http.StatusSeeOther, "/login")
}
