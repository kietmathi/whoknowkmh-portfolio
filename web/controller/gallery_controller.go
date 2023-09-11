package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/internal"
)

type galleryController struct {
	photoService domain.PhotoService
}

type GalleryController interface {
	ShowAll(c *gin.Context)
	ShowByID(c *gin.Context)
}

func NewPhotosController(ps domain.PhotoService) GalleryController {
	return &galleryController{
		photoService: ps,
	}
}

func (pc *galleryController) ShowAll(c *gin.Context) {
	var photos []domain.Photo
	photos, err := pc.photoService.FindAll()
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"gallery.all.html",
			gin.H{"error": "Error while getting photos"})
		return
	}
	c.HTML(
		http.StatusOK,
		"gallery.all.html",
		gin.H{
			"title":  "gallery",
			"photos": photos,
		},
	)
}

func (pc *galleryController) ShowByID(c *gin.Context) {
	imgIDParam := c.Param("imgid")
	imgID, err := strconv.Atoi(imgIDParam)
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"gallery.single.html",
			gin.H{"error": "Invalid photo id:" + imgIDParam})
		return
	}

	photo, err := pc.photoService.FindByID(uint(imgID))
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"gallery.single.html",
			gin.H{"error": "Error while getting photo id:" + imgIDParam})
		return
	}

	photoIds, err := pc.photoService.GetAllID()
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"gallery.single.html",
			gin.H{"error": "Error while getting photo id:" + imgIDParam})
		return
	}

	preID, nextID := internal.FindNextAndPrevPhotoID(photoIds, imgIDParam)
	data := make(map[string]interface{})
	data["title"] = "gallery"
	data["photo"] = photo
	data["description"] = template.HTML(photo.Description)
	data["preID"] = preID
	data["nextID"] = nextID
	c.HTML(
		http.StatusOK,
		"gallery.single.html",
		gin.H{
			"data": data,
		},
	)
}
