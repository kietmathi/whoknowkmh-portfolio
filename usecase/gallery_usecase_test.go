package usecase_test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/domain/mocks"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var mockPhotoRepoGu *mocks.PhotoRepository

func init() {
	// Khởi tạo mockPhotoRepository ở đây, một lần duy nhất
	mockPhotoRepoGu = new(mocks.PhotoRepository)
}

func TestGallryUsecaseFindPhotoByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test",
			Url:         "https://test.png",
			Description: "test",
			DeleteFlag:  newTrue(),
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// behavior
		mockPhotoRepoGu.
			On("FindByID", mock.AnythingOfType("uint")).
			Return(expectedPhoto, nil).Once()

		// mock usecase
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)
		photo, err := gu.FindPhotoByID(uint(1))

		// kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, expectedPhoto, photo)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoGu.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		emptyPhoto := domain.Photo{}

		// behavior
		mockPhotoRepoGu.
			On("FindByID", mock.AnythingOfType("uint")).
			Return(emptyPhoto, gorm.ErrRecordNotFound).Once() // Trả về lỗi gorm.ErrRecordNotFound

		// mock usecase
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)
		photo, err := gu.FindPhotoByID(uint(2))

		// kiểm tra kết quả
		assert.Error(t, err)
		assert.Equal(t, emptyPhoto, photo)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoGu.AssertExpectations(t)
	})
}

func TestGallryUsecaseFindAllAvailablePhoto(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedPhotos := []domain.Photo{
			{
				ID:          uint(2),
				Name:        "test",
				Url:         "https://test2.png",
				Description: "test",
				DeleteFlag:  newTrue(),
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				ID:          uint(1),
				Name:        "test",
				Url:         "https://test.png",
				Description: "test",
				DeleteFlag:  newTrue(),
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		}

		// behavior
		mockPhotoRepoGu.
			On("FindAllAvailable").
			Return(expectedPhotos, nil).Once()

		// mock usecase
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)
		photos, err := gu.FindAllAvailablePhoto()

		// kiểm tra kết quả
		assert.NoError(t, err)
		assert.Len(t, photos, len(expectedPhotos))
		assert.Equal(t, expectedPhotos, photos)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoGu.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedPhotos := []domain.Photo{}

		// behavior
		mockPhotoRepoGu.
			On("FindAllAvailable").
			Return(expectedPhotos, gorm.ErrRecordNotFound).Once()

		// mock usecase
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)
		photos, err := gu.FindAllAvailablePhoto()

		// kiểm tra kết quả
		assert.Error(t, err)
		assert.Len(t, photos, len(expectedPhotos))
		assert.Equal(t, expectedPhotos, photos)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoGu.AssertExpectations(t)
	})
}

func TestGallryUsecaseFindNextAndPrevPhotoID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		photoIds := []string{"photo1", "photo2", "photo3", "photo4"}

		// behavior cho repository
		mockPhotoRepoGu.
			On("GetAllAvailableID").
			Return(photoIds, nil).Once()

		// mock usecase
		targetID := "photo2"
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)
		prevID, nextID, err := gu.FindNextAndPrevPhotoID(targetID)

		// kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, "photo1", prevID)
		assert.Equal(t, "photo3", nextID)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoGu.AssertExpectations(t)
	})

	t.Run("success_2", func(t *testing.T) {
		photoIds := []string{"photo1", "photo2", "photo3", "photo4"}

		// behavior cho repository
		mockPhotoRepoGu.
			On("GetAllAvailableID").
			Return(photoIds, nil).Once()

		// mock usecase
		targetID := "photo1"
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)
		prevID, nextID, err := gu.FindNextAndPrevPhotoID(targetID)

		// kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, "", prevID)
		assert.Equal(t, "photo2", nextID)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoGu.AssertExpectations(t)
	})

	t.Run("success_3", func(t *testing.T) {
		photoIds := []string{"photo1", "photo2", "photo3", "photo4"}

		// behavior cho repository
		mockPhotoRepoGu.
			On("GetAllAvailableID").
			Return(photoIds, nil).Once()

		// mock usecase
		targetID := "photo4"
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)
		prevID, nextID, err := gu.FindNextAndPrevPhotoID(targetID)

		// mock usecase
		assert.NoError(t, err)
		assert.Equal(t, "photo3", prevID)
		assert.Equal(t, "", nextID)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoGu.AssertExpectations(t)
	})

	t.Run("success_4", func(t *testing.T) {
		photoIds := []string{"photo1", "photo2", "photo3", "photo4"}

		// behavior cho repository
		mockPhotoRepoGu.
			On("GetAllAvailableID").
			Return(photoIds, nil).Once()

		// mock usecase
		targetID := "photo6"
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)
		prevID, nextID, err := gu.FindNextAndPrevPhotoID(targetID)

		// kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, "", prevID)
		assert.Equal(t, "", nextID)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoGu.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		photoIds := []string{}

		// behavior cho repository
		mockPhotoRepoGu.
			On("GetAllAvailableID").
			Return(photoIds, gorm.ErrRecordNotFound).Once()

		// mock usecase
		targetID := "photo4"
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)
		prevID, nextID, err := gu.FindNextAndPrevPhotoID(targetID)

		// kiểm tra kết quả
		assert.Error(t, err)
		assert.Equal(t, "", prevID)
		assert.Equal(t, "", nextID)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoGu.AssertExpectations(t)
	})
}

func TestGallryUsecaseRenderTemplateWithGin(t *testing.T) {
	tpl := "template.html"

	t.Run("success", func(t *testing.T) {

		// Tạo một đối tượng use case mới
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)

		// Tạo một bộ kiểm tra HTTP và router Gin giả
		r := gin.Default()

		templ := template.Must(template.New(tpl).Parse(`Hello {{.data.title}}`))
		r.SetHTMLTemplate(templ)

		// Định nghĩa một điểm cuối API để kiểm tra use case
		r.GET("/render-template", func(c *gin.Context) {
			// Giả lập việc middleware đã xử lý CSRF và đã đặt giá trị cho csrf.TemplateTag

			data := make(map[string]interface{}, 1)
			data["title"] = domain.GalleryTitle
			gu.RenderTemplate(
				c,
				http.StatusOK,
				tpl,
				1*time.Minute,
				data)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/render-template", nil)
		r.ServeHTTP(w, req)

		// Kiểm tra xem phản ứng có đúng không
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Hello "+domain.GalleryTitle, w.Body.String())
		assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	})
	t.Run("error", func(t *testing.T) {

		// Tạo một đối tượng use case mới
		gu := usecase.NewGalleryUsecase(mockPhotoRepoGu)

		// Tạo một bộ kiểm tra HTTP và router Gin giả
		r := gin.Default()

		// Định nghĩa một điểm cuối API để kiểm tra use case
		r.GET("/render-template", func(c *gin.Context) {
			// render với template không tồn tại (chưa được set)
			data := make(map[string]interface{}, 1)
			data["title"] = domain.GalleryTitle
			gu.RenderTemplate(
				c,
				http.StatusOK,
				tpl,
				1*time.Minute,
				data)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/render-template", nil)
		r.ServeHTTP(w, req)

		// Kiểm tra kết quả
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
