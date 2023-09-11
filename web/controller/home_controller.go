package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

type homeController struct {
	photoService domain.PhotoService
}

type HomeController interface {
	Show(c *gin.Context)
}

func NewHomeController(ps domain.PhotoService) HomeController {
	return &homeController{
		photoService: ps,
	}
}

func (pc *homeController) Show(c *gin.Context) {
	photos, err := pc.photoService.FindAll()
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"home.html",
			gin.H{"error": "Error while getting photos"})
		return
	}

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"home.html",
		// Pass the data that the page uses
		gin.H{
			"title":  "home",
			"photos": photos,
		},
	)
}
