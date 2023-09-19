package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/bootstrap"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	errLogin, _ := lc.LoginUsecase.GetSession(c, "error")

	if errLogin == nil {
		errLogin = ""
	}

	lc.LoginUsecase.DeleteSession(c, "error")

	data := make(map[string]interface{})
	data["errLogin"] = errLogin
	lc.LoginUsecase.RenderTemplate(
		c,
		http.StatusOK,
		"admin/login.html",
		0*time.Second,
		data)
}

func (lgc *LoginController) LoginPost(c *gin.Context) {
	// Validate form input
	requestUser := domain.LoginUser{}
	if err := c.ShouldBind(&requestUser); err != nil {
		lgc.LoginUsecase.SetSession(c, "error", err.Error())
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	// Check for username and password match, usually from a database
	if requestUser.UserID != lgc.Env.AdminUserID ||
		bcrypt.CompareHashAndPassword([]byte(lgc.Env.AdminUserPwdHash), []byte(requestUser.Password)) != nil {
		lgc.LoginUsecase.SetSession(c, "error", "Invalid login credentials")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	accessToken, err := lgc.LoginUsecase.CreateAccessToken(&requestUser, lgc.Env.AccessTokenSecret, 2)
	if err != nil {
		lgc.LoginUsecase.SetSession(c, "error", err.Error())
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	lgc.LoginUsecase.DeleteSession(c, "error")

	lgc.LoginUsecase.SetCookieSession(c, "Authorization", accessToken, 2*3600)

	c.Redirect(http.StatusSeeOther, "/admin")
}
