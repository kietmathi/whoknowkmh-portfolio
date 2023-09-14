package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderutil"
)

type aboutController struct{}

type AboutController interface {
	Show(c *gin.Context)
}

func NewAboutController() AboutController {
	return &aboutController{}
}

func (ac *aboutController) Show(c *gin.Context) {
	isshow := c.DefaultQuery("isshow", "false")
	fmt.Println("isshow", isshow)
	data := make(map[string]interface{}, 1)
	data["title"] = "about"
	data["isshow"] = isshow
	renderutil.RenderTemplte(
		c,
		http.StatusOK,
		"about.html",
		1*time.Hour,
		data)
}
