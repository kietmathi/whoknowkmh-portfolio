package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type blogController struct{}

type BlogController interface {
	Show(c *gin.Context)
}

func NewBlogController() BlogController {
	return &blogController{}
}

func (bc *blogController) Show(c *gin.Context) {
	data := make(map[string]interface{})
	data["title"] = "about"
	c.HTML(
		http.StatusOK,
		"blog.html",
		gin.H{
			"data": data,
		},
	)
}
