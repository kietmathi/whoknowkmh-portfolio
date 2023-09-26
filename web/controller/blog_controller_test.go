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

func TestBlogController_Show(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		
		// mock usecase & controller
		mockBlogUsecase := new(mocks.BlogUsecase)
		blogController := &controller.BlogController{
			BlogUsecase: mockBlogUsecase}

		// behavior
		mockBlogUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			200,                                 // Kiểm tra giá trị statusCode
			domain.BlogTemplateName,             // Kiểm tra tên template
			1*time.Hour,                         // Kiểm tra cacheDuration
			mock.Anything,                       // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// Tạo router Gin
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/blog", blogController.Show)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/blog", nil)
		router.ServeHTTP(w, req)

		// Kiểm tra status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockBlogUsecase.AssertCalled(t,
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"),
			200,
			domain.BlogTemplateName,
			1*time.Hour,
			mock.Anything)
	})
}
