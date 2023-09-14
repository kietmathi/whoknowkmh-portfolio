package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderutil"
)

const templateTitle string = "gallery"

type galleryController struct {
	galleryService domain.GalleryService
	logger         domain.Logger
}

type GalleryController interface {
	ShowAll(c *gin.Context)
	ShowByID(c *gin.Context)
}

// NewGalleryController: create a new instance for GalleryController
func NewGalleryController(gs domain.GalleryService, l domain.Logger) GalleryController {
	return &galleryController{
		galleryService: gs,
		logger:         l,
	}
}

// ShowAll : When the user clicks on the 'Gallery' link,
// we should show the gallery page with all available photos
func (gc *galleryController) ShowAll(c *gin.Context) {
	templateName := "gallery.all.html"
	data := make(map[string]interface{}, 2)

	// Find all available photos in DB
	photos, err := gc.galleryService.FindAllPhoto()
	if err != nil {
		gc.logger.Printf("%+v\n", err)
		renderutil.RenderTemplte(
			c,
			http.StatusBadRequest,
			templateName,
			0*time.Second,
			data)
		return
	}

	// Rendering a gallery that shows all photos
	data["title"] = templateTitle
	data["photos"] = photos
	renderutil.RenderTemplte(
		c,
		http.StatusOK,
		templateName,
		1*time.Hour,
		data)
}

// ShowByID : When the user clicks on the link to a specific photo,
// we should show a page with relevant information about that photo
func (gc *galleryController) ShowByID(c *gin.Context) {
	templateName := "gallery.single.html"
	data := make(map[string]interface{}, 5)

	// Extract the photo ID from the request parameter
	imgIDParam := c.Param("imgid")
	imgID, err := strconv.Atoi(imgIDParam)
	if err != nil {
		gc.logger.Printf("%+v\n", err)
		renderutil.RenderTemplte(
			c,
			http.StatusBadRequest,
			templateName,
			0*time.Second,
			data)
		return
	}

	// Find photo information with a specific ID from the DB
	photo, err := gc.galleryService.FindPhotoByID(uint(imgID))
	if err != nil {
		gc.logger.Printf("%+v\n", err)
		renderutil.RenderTemplte(
			c,
			http.StatusBadRequest,
			templateName,
			0*time.Second,
			data)
		return
	}

	// find the next and previous photo IDs related to a specific ID
	// so that we can generate URLs for navigating to the adjacent photos.
	preID, nextID, err := gc.galleryService.FindNextAndPrevPhotoID(imgIDParam)
	if err != nil {
		gc.logger.Printf("%+v\n", err)
		renderutil.RenderTemplte(
			c,
			http.StatusBadRequest,
			templateName,
			0*time.Second,
			data)
		return
	}

	// rendering a page that shows information for a specific photo ID
	// and includes URLs for navigating to the adjacent photos.
	data["title"] = templateTitle
	data["photo"] = photo
	data["description"] = template.HTML(photo.Description)
	data["preID"] = preID
	data["nextID"] = nextID
	renderutil.RenderTemplte(
		c,
		http.StatusOK,
		templateName,
		1*time.Hour,
		data)
}
