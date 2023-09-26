package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

type GalleryController struct {
	GalleryUsecase domain.GalleryUsecase
	Logger         domain.Logger
}

// ShowAll : When the user clicks on the 'Gallery' link,
// we should show the gallery page with all available photos
func (gc *GalleryController) ShowAll(c *gin.Context) {
	data := make(map[string]interface{}, 2)

	// Find all available photos in DB
	photos, err := gc.GalleryUsecase.FindAllAvailablePhoto()
	if err != nil {
		gc.Logger.Printf("%+v\n", err)
		gc.GalleryUsecase.RenderTemplate(
			c,
			http.StatusBadRequest,
			domain.GalleryAllTemplateName,
			0*time.Second,
			data)
		return
	}

	// Rendering a gallery that shows all photos
	data["title"] = domain.GalleryTitle
	data["photos"] = photos
	gc.GalleryUsecase.RenderTemplate(
		c,
		http.StatusOK,
		domain.GalleryAllTemplateName,
		1*time.Hour,
		data)
}

// ShowByID : When the user clicks on the link to a specific photo,
// we should show a page with relevant information about that photo
func (gc *GalleryController) ShowByID(c *gin.Context) {
	data := make(map[string]interface{}, 5)

	// Extract the photo ID from the request parameter
	imgIDParam := c.Param("imgid")
	imgID, err := strconv.Atoi(imgIDParam)
	if err != nil {
		gc.Logger.Printf("%+v\n", err)
		gc.GalleryUsecase.RenderTemplate(
			c,
			http.StatusNotFound,
			"user/not_found.html",
			0*time.Second,
			data)
		return
	}

	// Find photo information with a specific ID from the DB
	photo, err := gc.GalleryUsecase.FindPhotoByID(uint(imgID))
	if err != nil {
		gc.Logger.Printf("%+v\n", err)
		gc.GalleryUsecase.RenderTemplate(
			c,
			http.StatusBadRequest,
			domain.GallerySingleTemplateName,
			0*time.Second,
			data)
		return
	}

	// find the next and previous photo IDs related to a specific ID
	// so that we can generate URLs for navigating to the adjacent photos.
	preID, nextID, err := gc.GalleryUsecase.FindNextAndPrevPhotoID(imgIDParam)
	if err != nil {
		gc.Logger.Printf("%+v\n", err)
		gc.GalleryUsecase.RenderTemplate(
			c,
			http.StatusBadRequest,
			domain.GallerySingleTemplateName,
			0*time.Second,
			data)
		return
	}

	// rendering a page that shows information for a specific photo ID
	// and includes URLs for navigating to the adjacent photos.
	data["title"] = domain.GalleryTitle
	data["photo"] = photo
	data["description"] = template.HTML(photo.Description)
	data["preID"] = preID
	data["nextID"] = nextID
	gc.GalleryUsecase.RenderTemplate(
		c,
		http.StatusOK,
		domain.GallerySingleTemplateName,
		1*time.Hour,
		data)
}
