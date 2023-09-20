package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
)

type AdminController struct {
	AdminUsecase domain.AdminUsecase
	Logger       domain.Logger
}

func (ac *AdminController) Show(c *gin.Context) {
	tableNames := ac.AdminUsecase.FindAvailableDBTable()
	data := make(map[string]interface{}, 2)
	data["title"] = domain.AdminTitle
	data["tableNames"] = tableNames
	ac.AdminUsecase.RenderTemplate(
		c,
		http.StatusOK,
		domain.AdminTemplateName,
		0*time.Second,
		data,
	)
}

func (ac *AdminController) ShowTableDetail(c *gin.Context) {
	tableName := c.Param("tablename")

	var data interface{}
	var err error
	switch tableName {
	case "photo":
		data, err = ac.AdminUsecase.ShowAllPhoto()
	}

	if err != nil {
		ac.Logger.Printf("%+v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (ac *AdminController) UpdatePhoto(c *gin.Context) {
	photo := domain.Photo{}
	if err := c.ShouldBind(&photo); err != nil {
		ac.Logger.Printf("%+v\n", err)
		// Xử lý lỗi nếu có
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	updatedPhoto, err := ac.AdminUsecase.UpdatePhotoByID(photo)
	if err != nil {
		ac.Logger.Printf("%+v\n", err)
		// Xử lý lỗi nếu có
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPhoto)
}

func (ac *AdminController) InsertPhoto(c *gin.Context) {
	photo := domain.Photo{}
	if err := c.ShouldBind(&photo); err != nil {
		ac.Logger.Printf("%+v\n", err)
		// Xử lý lỗi nếu có
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	insertPhoto, err := ac.AdminUsecase.InsertPhoto(photo)
	if err != nil {
		ac.Logger.Printf("%+v\n", err)
		// Xử lý lỗi nếu có
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, insertPhoto)
}
