package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderutil"
)

type aboutController struct{}

type AboutController interface {
	Show(c *gin.Context)
}

// NewAboutController: create a new instance for AboutController
func NewAboutController() AboutController {
	return &aboutController{}
}

// Show: When the user clicks on the 'About' link, we should show the about page with relevant information
func (ac *aboutController) Show(c *gin.Context) {
	isshow := c.DefaultQuery("isshow", "false")
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
