package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderutil"
)

type homeController struct{}

type HomeController interface {
	Show(c *gin.Context)
}

func NewHomeController() HomeController {
	return &homeController{}
}

func (pc *homeController) Show(c *gin.Context) {
	data := make(map[string]interface{}, 1)
	// Call the HTML method of the Context to render a template
	data["title"] = "home"
	renderutil.RenderTemplte(
		c,
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the home.html template
		"home.html",
		// Set cache time
		1*time.Hour,
		// Pass the data that the page uses
		data)
}
