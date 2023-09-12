package controller

import (
	"html/template"
	"net/http"
	"strconv"

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

func NewGalleryController(gs domain.GalleryService, l domain.Logger) GalleryController {
	return &galleryController{
		galleryService: gs,
		logger:         l,
	}
}

func (gc *galleryController) ShowAll(c *gin.Context) {
	templateName := "gallery.all.html"
	data := make(map[string]interface{}, 2)

	photos, err := gc.galleryService.FindAllPhoto()
	if err != nil {
		gc.logger.Printf("%+v\n", err)
		renderutil.RenderTemplte(
			c,
			http.StatusBadRequest,
			templateName,
			data)
		return
	}

	data["title"] = templateTitle
	data["photos"] = photos
	renderutil.RenderTemplte(
		c,
		http.StatusOK,
		templateName,
		data)
}

func (gc *galleryController) ShowByID(c *gin.Context) {
	templateName := "gallery.single.html"
	data := make(map[string]interface{}, 5)

	imgIDParam := c.Param("imgid")
	imgID, err := strconv.Atoi(imgIDParam)
	if err != nil {
		gc.logger.Printf("%+v\n", err)
		renderutil.RenderTemplte(
			c,
			http.StatusBadRequest,
			templateName,
			data)
		return
	}

	photo, err := gc.galleryService.FindPhotoByID(uint(imgID))
	if err != nil {
		gc.logger.Printf("%+v\n", err)
		renderutil.RenderTemplte(
			c,
			http.StatusBadRequest,
			templateName,
			data)
		return
	}

	preID, nextID, err := gc.galleryService.FindNextAndPrevPhotoID(imgIDParam)
	if err != nil {
		gc.logger.Printf("%+v\n", err)
		renderutil.RenderTemplte(
			c,
			http.StatusBadRequest,
			templateName,
			data)
		return
	}

	data["title"] = templateTitle
	data["photo"] = photo
	data["description"] = template.HTML(photo.Description)
	data["preID"] = preID
	data["nextID"] = nextID
	renderutil.RenderTemplte(
		c,
		http.StatusOK,
		templateName,
		data)
}
