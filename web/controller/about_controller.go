package controller

import (
	"net/http"

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
	data := make(map[string]interface{}, 1)
	renderutil.RenderTemplte(
		c,
		http.StatusOK,
		"about.html",
		data)
}
