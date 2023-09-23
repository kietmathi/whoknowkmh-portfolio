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

var mockPhotoRepoAu *mocks.PhotoRepository

func init() {
	// Khởi tạo mockPhotoRepository ở đây, một lần duy nhất
	mockPhotoRepoAu = new(mocks.PhotoRepository)
}

var now = time.Date(2023, time.September, 20, 22, 33, 49, 201879700, time.UTC)

func newTrue() *bool {
	t := true
	return &t
}

func TestAdminUsecaseFindAvailableDBTable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		photoTableName := "photo"
		expectedTableNames := make([]string, 0)
		expectedTableNames = append(expectedTableNames, photoTableName)

		// behavior
		mockPhotoRepoAu.
			On("TableName").
			Return(photoTableName).Once()

		// mock usecase
		au := usecase.NewAdminUsecase(mockPhotoRepoAu)
		tableNames := au.FindAvailableDBTable()

		// kiểm tra kết quả
		assert.Equal(t, expectedTableNames, tableNames)
		assert.Len(t, tableNames, len(expectedTableNames))

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoAu.AssertExpectations(t)
	})
}

func TestAdminUsecaseShowAllPhoto(t *testing.T) {
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
		mockPhotoRepoAu.
			On("FindAll").
			Return(expectedPhotos, nil).
			Once()

		// mock usecase
		au := usecase.NewAdminUsecase(mockPhotoRepoAu)
		photos, err := au.ShowAllPhoto()

		// kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, expectedPhotos, photos)
		assert.Len(t, photos, len(expectedPhotos))

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoAu.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedPhotos := []domain.Photo{}

		// behavior
		mockPhotoRepoAu.
			On("FindAll").
			Return(expectedPhotos, gorm.ErrRecordNotFound).
			Once()

		// mock usecase
		au := usecase.NewAdminUsecase(mockPhotoRepoAu)
		photos, err := au.ShowAllPhoto()

		// kiểm tra kết quả
		assert.Error(t, err)
		assert.Equal(t, expectedPhotos, photos)
		assert.Len(t, photos, len(expectedPhotos))

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoAu.AssertExpectations(t)
	})
}

func TestAdminUsecaseUpdatePhotoByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		preUpdatePhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test update",
			Url:         "https://test.png",
			Description: "test",
			DeleteFlag:  newTrue(),
			CreatedAt:   now,
		}

		expectedPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test update",
			Url:         "https://test.png",
			Description: "test",
			DeleteFlag:  newTrue(),
			CreatedAt:   now,
			UpdatedAt:   time.Now(),
		}

		// behavior
		mockPhotoRepoAu.
			On("UpdateByID", mock.AnythingOfType("domain.Photo")).
			Return(expectedPhoto, nil).
			Once()

		// mock usecase
		au := usecase.NewAdminUsecase(mockPhotoRepoAu)
		updatedPhoto, err := au.UpdatePhotoByID(preUpdatePhoto)

		// kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, expectedPhoto, updatedPhoto)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoAu.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		preUpdatePhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test update",
			Url:         "https://test.png",
			Description: "test",
			DeleteFlag:  newTrue(),
			CreatedAt:   now,
		}

		// behavior
		mockPhotoRepoAu.
			On("UpdateByID", mock.AnythingOfType("domain.Photo")).
			Return(preUpdatePhoto, gorm.ErrRecordNotFound).
			Once()

		// mock usecase
		au := usecase.NewAdminUsecase(mockPhotoRepoAu)
		updatedPhoto, err := au.UpdatePhotoByID(preUpdatePhoto)

		// kiểm tra kết quả
		assert.Error(t, err)
		assert.Equal(t, preUpdatePhoto, updatedPhoto)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoAu.AssertExpectations(t)
	})
}

func TestAdminUsecaseInsertPhoto(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		photo := domain.Photo{
			Name:        "test update",
			Url:         "https://test.png",
			Description: "test",
		}

		expectedPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test update",
			Url:         "https://test.png",
			Description: "test",
			DeleteFlag:  newTrue(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// behavior
		mockPhotoRepoAu.
			On("InsertPhoto", mock.AnythingOfType("domain.Photo")).
			Return(expectedPhoto, nil).
			Once()

		// mock usecase
		au := usecase.NewAdminUsecase(mockPhotoRepoAu)
		insertedPhoto, err := au.InsertPhoto(photo)

		// kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, expectedPhoto, insertedPhoto)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockPhotoRepoAu.AssertExpectations(t)
	})
}

func TestAdminUsecaseRenderTemplateWithGin(t *testing.T) {
	tpl := "template.html"

	t.Run("success", func(t *testing.T) {

		// Tạo một đối tượng use case
		gu := usecase.NewAdminUsecase(mockPhotoRepoAu)

		// Tạo một bộ kiểm tra HTTP và router Gin
		r := gin.Default()

		// set template
		templ := template.Must(template.New(tpl).Parse(`Hello {{.data.title}}`))
		r.SetHTMLTemplate(templ)

		// Định nghĩa một endpoint API để kiểm tra use case
		r.GET("/render-template", func(c *gin.Context) {
			// render template
			data := make(map[string]interface{}, 1)
			data["title"] = domain.AdminTitle
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
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Hello "+domain.AdminTitle, w.Body.String())
		assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	})
	t.Run("error", func(t *testing.T) {

		// Tạo một đối tượng use case mới
		gu := usecase.NewAdminUsecase(mockPhotoRepoAu)

		// Tạo một bộ kiểm tra HTTP và router Gin
		r := gin.Default()

		// Định nghĩa một endpoint API để kiểm tra use case
		r.GET("/render-template", func(c *gin.Context) {

			// render với template không tồn tại (chưa được set
			data := make(map[string]interface{}, 1)
			data["title"] = domain.AdminTitle
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
