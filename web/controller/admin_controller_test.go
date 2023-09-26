package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
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
	"gorm.io/gorm"
)

func TestAdminController_HandleError(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		// mock usecase, logger & controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		mockLogger := new(mocks.Logger)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase,
			Logger:       mockLogger,
		}

		// behavior
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		mockErr := errors.New("test error")
		// Tạo một router Gin
		router := gin.Default()
		router.GET("/admin/handle-error", func(c *gin.Context) {
			// Gọi hàm handleError với một mock error
			adminController.HandleError(c, mockErr)
		})

		// expected response body
		body, err := json.Marshal(domain.ErrorResponse{Message: mockErr.Error()})
		assert.NoError(t, err)
		bodyString := string(body)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/admin/handle-error", nil)
		router.ServeHTTP(w, req)

		// Kiểm tra kết quả
		assert.Equal(t, http.StatusBadRequest, w.Code) // Xác minh rằng mã lỗi HTTP đã được đặt đúng
		assert.Equal(t, bodyString, w.Body.String())   // kiểm tra response body

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockLogger.AssertExpectations(t)
	})
}

func TestAdminController_Show(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		// Tạo router Gin
		router := gin.Default()

		// mock usecase & controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase}

		// behavior

		expectedTableNames := []string{"photo"}
		mockAdminUsecase.On("FindAvailableDBTable").Return(expectedTableNames)

		mockAdminUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			200,                                 // Kiểm tra giá trị statusCode
			domain.AdminTemplateName,            // Kiểm tra tên template
			0*time.Hour,                         // Kiểm tra cacheDuration
			mock.Anything,                       // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// Định nghĩa endpoint API và gọi controller method
		router.GET("/admin", adminController.Show)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/admin", nil)
		router.ServeHTTP(w, req)

		// Kiểm tra response có đúng mã trạng thái
		assert.Equal(t, http.StatusOK, w.Code)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAdminUsecase.AssertExpectations(t)
	})
}

func TestBlogController_ShowTablePhoto(t *testing.T) {
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		mockPhotos := []domain.Photo{
			{
				ID:          uint(2),
				Name:        "test",
				Url:         "https://test2.png",
				Description: "test",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				ID:          uint(1),
				Name:        "test",
				Url:         "https://test.png",
				Description: "test",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		}

		// mock usecase, logger & controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		mockLogger := new(mocks.Logger)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase,
			Logger:       mockLogger}

		// behavior
		mockAdminUsecase.On("ShowAllPhoto").Return(mockPhotos, nil)

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/admin/photo", adminController.ShowTablePhoto)

		// mock response body
		body, err := json.Marshal(mockPhotos)
		assert.NoError(t, err)
		bodyString := string(body)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/admin/photo", nil)
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusOK, w.Code)
		// kiểm tra response body
		assert.Equal(t, bodyString, w.Body.String())

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAdminUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockPhotos := []domain.Photo{}

		// mock usecase, logger && controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		mockLogger := new(mocks.Logger)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase,
			Logger:       mockLogger}

		// behavior
		mockErr := gorm.ErrRecordNotFound
		mockAdminUsecase.On("ShowAllPhoto").Return(mockPhotos, mockErr)
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/admin/photo", adminController.ShowTablePhoto)

		// mock response body
		body, err := json.Marshal(domain.ErrorResponse{Message: mockErr.Error()})
		assert.NoError(t, err)
		bodyString := string(body)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/admin/photo", nil)
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// kiểm tra response body
		assert.Equal(t, bodyString, w.Body.String())

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAdminUsecase.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}

func TestBlogController_UpdatePhoto(t *testing.T) {
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		mockRequestPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test",
			Url:         "https://test2.png",
			Description: "test",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		mockUpdatedPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test update",
			Url:         "https://test2.png",
			Description: "test",
			CreatedAt:   now,
			UpdatedAt:   time.Now(),
		}

		// mock usecase, logger & controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		mockLogger := new(mocks.Logger)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase,
			Logger:       mockLogger}

		// behavior
		mockAdminUsecase.On("UpdatePhotoByID", mock.AnythingOfType("domain.Photo")).Return(mockUpdatedPhoto, nil)

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.PUT("/admin/photo/:id", adminController.UpdatePhoto)

		// mock request photo
		jsonPhoto, err := json.Marshal(mockRequestPhoto)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("PUT", "/admin/photo/1", bytes.NewBuffer(jsonPhoto))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// mock response body
		body, err := json.Marshal(mockUpdatedPhoto)
		assert.NoError(t, err)
		bodyString := string(body)

		// kiểm tra statuc code
		assert.Equal(t, http.StatusOK, w.Code)
		// kiểm tra response body
		assert.Equal(t, bodyString, w.Body.String())

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAdminUsecase.AssertExpectations(t)
	})

	t.Run("error_bind", func(t *testing.T) {
		mockRequestPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "",
			Url:         "https://test2.png",
			Description: "test",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// mock usecase, logger & controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		mockLogger := new(mocks.Logger)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase,
			Logger:       mockLogger}

		// behavior
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.PUT("/admin/photo/:id", adminController.UpdatePhoto)

		// mock request photo
		jsonPhoto, err := json.Marshal(mockRequestPhoto)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		// goi endpoint và kiểm tra repsonse
		req, _ := http.NewRequest("PUT", "/admin/photo/1", bytes.NewBuffer(jsonPhoto))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// mock response body
		mockErr := errors.New("Key: 'Photo.Name' Error:Field validation for 'Name' failed on the 'required' tag")
		body, err := json.Marshal(domain.ErrorResponse{Message: mockErr.Error()})
		assert.NoError(t, err)
		bodyString := string(body)

		// kiểm tra status code
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// kiểm tra response body
		assert.Equal(t, bodyString, w.Body.String())

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAdminUsecase.AssertExpectations(t)
	})

	t.Run("error_update_photo", func(t *testing.T) {
		mockRequestPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test",
			Url:         "https://test2.png",
			Description: "test",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// mock usecase, logger & controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		mockLogger := new(mocks.Logger)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase,
			Logger:       mockLogger}

		// behavior
		mockErr := gorm.ErrRecordNotFound
		mockAdminUsecase.On("UpdatePhotoByID", mock.AnythingOfType("domain.Photo")).Return(mockRequestPhoto, mockErr)
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.PUT("/admin/photo/:id", adminController.UpdatePhoto)

		// mock repsponse body
		body, err := json.Marshal(domain.ErrorResponse{Message: mockErr.Error()})
		assert.NoError(t, err)
		bodyString := string(body)

		// mock request photo
		jsonPhoto, err := json.Marshal(mockRequestPhoto)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("PUT", "/admin/photo/1", bytes.NewBuffer(jsonPhoto))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// kiểm tra response body
		assert.Equal(t, bodyString, w.Body.String())

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAdminUsecase.AssertExpectations(t)
	})
}

func TestBlogController_InsertPhoto(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockRequestPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test",
			Url:         "https://test2.png",
			Description: "test",
		}

		// mock usecase, logger & controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		mockLogger := new(mocks.Logger)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase,
			Logger:       mockLogger}

		// behavior
		mockAdminUsecase.On("InsertPhoto", mock.AnythingOfType("domain.Photo")).Return(mockRequestPhoto, nil)

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.POST("/admin/photo/:id", adminController.InsertPhoto)

		// mock request photo
		jsonPhoto, err := json.Marshal(mockRequestPhoto)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("POST", "/admin/photo/1", bytes.NewBuffer(jsonPhoto))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// mocke response body
		body, err := json.Marshal(mockRequestPhoto)
		assert.NoError(t, err)
		bodyString := string(body)

		// kiểm tra status code
		assert.Equal(t, http.StatusOK, w.Code)
		// kiểm tra response body
		assert.Equal(t, bodyString, w.Body.String())

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAdminUsecase.AssertExpectations(t)
	})

	t.Run("error_bind", func(t *testing.T) {
		mockRequestPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "",
			Url:         "https://test2.png",
			Description: "test",
		}

		// mock usecase, logger && controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		mockLogger := new(mocks.Logger)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase,
			Logger:       mockLogger}

		// behavior
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.POST("/admin/photo/:id", adminController.InsertPhoto)

		// mock request photo
		jsonPhoto, err := json.Marshal(mockRequestPhoto)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response body
		req, _ := http.NewRequest("POST", "/admin/photo/1", bytes.NewBuffer(jsonPhoto))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// mock response body
		mockErr := errors.New("Key: 'Photo.Name' Error:Field validation for 'Name' failed on the 'required' tag")
		body, err := json.Marshal(domain.ErrorResponse{Message: mockErr.Error()})
		assert.NoError(t, err)
		bodyString := string(body)

		// kiểm tra status code
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// kiểm tra response body
		assert.Equal(t, bodyString, w.Body.String())

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAdminUsecase.AssertExpectations(t)
	})

	t.Run("error_insert_photo", func(t *testing.T) {
		mockRequestPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test",
			Url:         "https://test2.png",
			Description: "test",
		}

		// mock usecase, logger & controller
		mockAdminUsecase := new(mocks.AdminUsecase)
		mockLogger := new(mocks.Logger)
		adminController := &controller.AdminController{
			AdminUsecase: mockAdminUsecase,
			Logger:       mockLogger}

		// behavior
		mockErr := gorm.ErrRecordNotFound
		mockAdminUsecase.On("InsertPhoto", mock.AnythingOfType("domain.Photo")).Return(mockRequestPhoto, mockErr)
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.POST("/admin/photo/:id", adminController.InsertPhoto)

		// mock response body
		body, err := json.Marshal(domain.ErrorResponse{Message: mockErr.Error()})
		assert.NoError(t, err)
		bodyString := string(body)

		// mock request photo
		jsonPhoto, err := json.Marshal(mockRequestPhoto)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response body
		req, _ := http.NewRequest("POST", "/admin/photo/1", bytes.NewBuffer(jsonPhoto))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// kiểm tra response body
		assert.Equal(t, bodyString, w.Body.String())

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockAdminUsecase.AssertExpectations(t)
	})
}
