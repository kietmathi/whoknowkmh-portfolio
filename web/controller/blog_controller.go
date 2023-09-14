package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/internal/renderutil"
)

type blogController struct{}

type BlogController interface {
	Show(c *gin.Context)
}

func NewBlogController() BlogController {
	return &blogController{}
}

func (bc *blogController) Show(c *gin.Context) {
	data := make(map[string]interface{}, 1)
	data["title"] = "blog"
	renderutil.RenderTemplte(
		c,
		http.StatusOK,
		"blog.html",
		1*time.Hour,
		data)
}
