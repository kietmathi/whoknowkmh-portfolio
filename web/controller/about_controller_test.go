package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/domain/mocks"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAboutController_Show(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		// Tạo router Gin
		router := gin.Default()

		// mock usecase & controller
		mockAboutUsecase := new(mocks.AboutUsecase)
		aboutController := &controller.AboutController{
			AboutUsecase: mockAboutUsecase}

		// behavior
		mockAboutUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			http.StatusOK,                       // Kiểm tra giá trị statusCode
			domain.AboutTemplateName,            // Kiểm tra tên template
			1*time.Hour,                         // Kiểm tra cacheDuration
			mock.Anything,                       // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// Định nghĩa endpoint API và gọi controller method
		router.GET("/about", aboutController.Show)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/about", nil)
		router.ServeHTTP(w, req)

		// Kiểm tra status
		assert.Equal(t, http.StatusOK, w.Code)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAboutUsecase.AssertCalled(t,
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"),
			200,
			domain.AboutTemplateName,
			1*time.Hour,
			mock.Anything)
	})
}
