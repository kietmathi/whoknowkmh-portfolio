package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type aboutController struct{}

type AboutController interface {
	Show(c *gin.Context)
}

func NewAboutController() AboutController {
	return &aboutController{}
}

func (ac *aboutController) Show(c *gin.Context) {
	data := make(map[string]interface{})
	data["title"] = "about"
	c.HTML(
		http.StatusOK,
		"about.html",
		gin.H{
			"data": data,
		},
	)
}
