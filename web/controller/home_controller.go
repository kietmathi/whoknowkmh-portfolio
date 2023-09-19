package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

type HomeController struct {
	HomeUsecase domain.HomeUsecase
}

// Show: When the user clicks on the 'Home' link, we should show the home page with relevant information
func (hc *HomeController) Show(c *gin.Context) {
	data := make(map[string]interface{}, 1)
	// Call the HTML method of the Context to render a template
	data["title"] = "home"
	hc.HomeUsecase.RenderTemplate(
		c,
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the home.html template
		"user/home.html",
		// Set cache time
		1*time.Hour,
		// Pass the data that the page uses
		data)
}
