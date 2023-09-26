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
	"gorm.io/gorm"
)

func TestGalleryControllerShow_All(t *testing.T) {
	now := time.Now()
	t.Run("success", func(t *testing.T) {

		// mock usecase, logger & controller
		mockGalleryUsecase := new(mocks.GalleryUsecase)
		mockLogger := new(mocks.Logger)
		galleryController := &controller.GalleryController{
			GalleryUsecase: mockGalleryUsecase,
			Logger:         mockLogger,
		}

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

		// behavior
		mockGalleryUsecase.On("FindAllAvailablePhoto").Return(mockPhotos, nil)
		mockGalleryUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			http.StatusOK,                       // Kiểm tra giá trị statusCode
			domain.GalleryAllTemplateName,       // Kiểm tra tên template
			1*time.Hour,                         // Kiểm tra cacheDuration
			mock.Anything,                       // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// Tạo router Gin
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/gallery", galleryController.ShowAll)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/gallery", nil)
		router.ServeHTTP(w, req)

		// Kiểm tra status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockGalleryUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {

		// mock usecase, logger & controller
		mockGalleryUsecase := new(mocks.GalleryUsecase)
		mockLogger := new(mocks.Logger)
		galleryController := &controller.GalleryController{
			GalleryUsecase: mockGalleryUsecase,
			Logger:         mockLogger,
		}

		// behavior
		mockPhotos := []domain.Photo{}
		mockGalleryUsecase.On("FindAllAvailablePhoto").Return(mockPhotos, gorm.ErrRecordNotFound)

		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		mockGalleryUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			http.StatusBadRequest,
			domain.GalleryAllTemplateName,
			0*time.Second, // Kiểm tra cacheDuration
			mock.Anything, // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// Tạo router Gin
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/gallery", galleryController.ShowAll)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/gallery", nil)
		router.ServeHTTP(w, req)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockGalleryUsecase.AssertExpectations(t)
	})
}

func TestGalleryController_ShowByID(t *testing.T) {
	now := time.Now()
	t.Run("success", func(t *testing.T) {

		// mock usecase, logger & controller
		mockGalleryUsecase := new(mocks.GalleryUsecase)
		mockLogger := new(mocks.Logger)
		galleryController := &controller.GalleryController{
			GalleryUsecase: mockGalleryUsecase,
			Logger:         mockLogger,
		}

		// behavior
		mockPhoto := domain.Photo{
			ID:          uint(2),
			Name:        "test",
			Url:         "https://test.png",
			Description: "test",
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		mockGalleryUsecase.On("FindPhotoByID", mock.AnythingOfType("uint")).Return(mockPhoto, nil)

		ids := []string{"1", "2", "3"}
		mockGalleryUsecase.On("FindNextAndPrevPhotoID", mock.AnythingOfType("string")).Return(ids[0], ids[2], nil)

		mockGalleryUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			http.StatusOK,                       // Kiểm tra giá trị statusCode
			domain.GallerySingleTemplateName,    // Kiểm tra tên template
			1*time.Hour,                         // Kiểm tra cacheDuration
			mock.Anything,                       // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// Tạo router Gin
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/gallery/:imgid", galleryController.ShowByID)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/gallery/2", nil)
		router.ServeHTTP(w, req)

		// Kiểm tra response có đúng mã trạng thái và dữ liệu mong đợi
		assert.Equal(t, http.StatusOK, w.Code)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockGalleryUsecase.AssertExpectations(t)
	})

	t.Run("error_parse_int", func(t *testing.T) {

		// mock usecase, logger & controller
		mockGalleryUsecase := new(mocks.GalleryUsecase)
		mockLogger := new(mocks.Logger)
		galleryController := &controller.GalleryController{
			GalleryUsecase: mockGalleryUsecase,
			Logger:         mockLogger,
		}

		// behavior
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		mockGalleryUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			http.StatusNotFound,
			"user/not_found.html",
			0*time.Second, // Kiểm tra cacheDuration
			mock.Anything, // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// Tạo router Gin
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/gallery/:imgid", galleryController.ShowByID)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/gallery/two", nil)
		router.ServeHTTP(w, req)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockGalleryUsecase.AssertExpectations(t)
	})

	t.Run("errorPhotoNotFound", func(t *testing.T) {

		// mock usecase, logger & controller
		mockGalleryUsecase := new(mocks.GalleryUsecase)
		mockLogger := new(mocks.Logger)
		galleryController := &controller.GalleryController{
			GalleryUsecase: mockGalleryUsecase,
			Logger:         mockLogger,
		}

		// bhavior
		mockPhoto := domain.Photo{}
		mockGalleryUsecase.On("FindPhotoByID", mock.AnythingOfType("uint")).Return(mockPhoto, gorm.ErrRecordNotFound)

		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		mockGalleryUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			http.StatusBadRequest,               // Kiểm tra giá trị statusCode
			domain.GallerySingleTemplateName,    // Kiểm tra tên template
			0*time.Hour,                         // Kiểm tra cacheDuration
			mock.Anything,                       // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// Tạo router Gin
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/gallery/:imgid", galleryController.ShowByID)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/gallery/2", nil)
		router.ServeHTTP(w, req)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockGalleryUsecase.AssertExpectations(t)
	})

	t.Run("errorFindNextAndPrevPhotoID", func(t *testing.T) {

		// mock usecase, logger & controller
		mockGalleryUsecase := new(mocks.GalleryUsecase)
		mockLogger := new(mocks.Logger)
		galleryController := &controller.GalleryController{
			GalleryUsecase: mockGalleryUsecase,
			Logger:         mockLogger,
		}

		// behavior

		mockPhoto := domain.Photo{
			ID:          uint(2),
			Name:        "test",
			Url:         "https://test.png",
			Description: "test",
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		mockGalleryUsecase.On("FindPhotoByID", mock.AnythingOfType("uint")).Return(mockPhoto, nil)

		mockGalleryUsecase.On("FindNextAndPrevPhotoID", mock.AnythingOfType("string")).Return("", "", gorm.ErrRecordNotFound)

		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()

		mockGalleryUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			http.StatusBadRequest,               // Kiểm tra giá trị statusCode
			domain.GallerySingleTemplateName,    // Kiểm tra tên template
			0*time.Hour,                         // Kiểm tra cacheDuration
			mock.Anything,                       // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// Tạo router Gin
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/gallery/:imgid", galleryController.ShowByID)

		w := httptest.NewRecorder()
		// Gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/gallery/2", nil)
		router.ServeHTTP(w, req)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockGalleryUsecase.AssertExpectations(t)
	})
}
