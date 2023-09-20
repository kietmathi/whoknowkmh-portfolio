package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
}

// Show: When the user clicks on the 'Blog' link, we should show the blog page with relevant information
func (bc *BlogController) Show(c *gin.Context) {
	data := make(map[string]interface{}, 1)
	data["title"] = domain.BlogTitle
	bc.BlogUsecase.RenderTemplate(
		c,
		http.StatusOK,
		domain.BlogTemplateName,
		1*time.Hour,
		data)
}
