package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

type AboutController struct {
	AboutUsecase domain.AboutUsecase
}

// Show: When the user clicks on the 'About' link, we should show the about page with relevant information
func (abc *AboutController) Show(c *gin.Context) {
	isshow := c.DefaultQuery("isshow", "false")
	data := make(map[string]interface{}, 1)
	data["title"] = domain.AboutTitle
	data["isshow"] = isshow
	abc.AboutUsecase.RenderTemplate(
		c,
		http.StatusOK,
		domain.AboutTemplateName,
		1*time.Hour,
		data)
}
