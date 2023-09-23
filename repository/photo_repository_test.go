package repository_test

import (
	"errors"
	"testing"
	"time"

	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/repository"
	"github.com/kietmathi/whoknowkmh-portfolio/sqlite/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var now = time.Date(2023, time.September, 20, 22, 33, 49, 201879700, time.UTC)

func newTrue() *bool {
	t := true
	return &t
}
func newFalse() *bool {
	f := false
	return &f
}

func TestTableName(t *testing.T) {
	databaseHelper := &mocks.Database{}

	// Thiết lập behavior cho mock
	expectedTableName := "photo"

	t.Run("success", func(t *testing.T) {
		// Tạo một mock cho repository
		pr := repository.NewPhotoRepository(databaseHelper)

		tableName := pr.TableName()

		// Kiểm tra kết quả
		assert.Equal(t, expectedTableName, tableName)
	})
}

func TestFindByID(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		// Tạo một mock cho databse
		databaseHelper := &mocks.Database{}

		expectedPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test",
			Url:         "https://test.png",
			Description: "test",
			DeleteFlag:  newFalse(),
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// Tạo một mock cho repository
		pr := repository.NewPhotoRepository(databaseHelper)

		photoPtr := &expectedPhoto

		// Thiết lập behavior cho mock
		// dựa theo hành vi dưới tại FindByID
		// p.DB.Model(&domain.Photo{}).First(&photo, fmt.Sprintf("id=%v", id))

		// p.DB.Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper) // Trả về chính databaseHelper

		// .First(&photo, fmt.Sprintf("id=%v", id))
		databaseHelper.
			On("First", mock.AnythingOfType("*domain.Photo"), mock.AnythingOfType("string")).
			// set giá trị của expectedPhoto cho đối số "photo" tại First(&photo, fmt.Sprintf("id=%v", id))
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*domain.Photo)
				*out = *photoPtr
			}).
			Return(nil) // Trả về nil để cho biết không có lỗi

		// Sử dụng mock trong unit test
		photo, err := pr.FindByID(uint(1))

		// Kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, expectedPhoto, photo)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một mock cho database
		databaseHelper := &mocks.Database{}

		expectedPhoto := domain.Photo{}

		photoPtr := &expectedPhoto

		// Thiết lập behavior cho mock
		// dựa theo hành vi dưới tại FindByID
		// p.DB.Model(&domain.Photo{}).First(&photo, fmt.Sprintf("id=%v", id))

		// p.DB.Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper) // Trả về chính databaseHelper

		// .First(&photo, fmt.Sprintf("id=%v", id))
		databaseHelper.
			On("First", mock.AnythingOfType("*domain.Photo"), mock.AnythingOfType("string")).
			// set giá trị của expectedPhoto cho đối số "photo" tại First(&photo, fmt.Sprintf("id=%v", id))
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*domain.Photo)
				*out = *photoPtr
			}).
			Return(gorm.ErrRecordNotFound) // Trả về nil để cho biết không có lỗi

		// Sử dụng mock trong unit test
		pr := repository.NewPhotoRepository(databaseHelper)
		photo, err := pr.FindByID(uint(1))

		// Kiểm tra kết quả
		assert.Error(t, err) // Kiểm tra xem có lỗi trả về hay không
		assert.Equal(t, expectedPhoto, photo)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})
}

func TestFindAllAvailable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Tạo một mock cho database
		databaseHelper := &mocks.Database{}

		expectedPhotos := []domain.Photo{
			{
				ID:          uint(2),
				Name:        "test",
				Url:         "https://test2.png",
				Description: "test",
				DeleteFlag:  newFalse(),
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				ID:          uint(1),
				Name:        "test",
				Url:         "https://test.png",
				Description: "test",
				DeleteFlag:  newFalse(),
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		}

		photosPtr := &expectedPhotos

		// behavior
		// p.db.Model(&domain.Photo{}).Where("delete_flag = ?", false).Order("id DESC").Find(&photos)

		// Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper)
		// Where("delete_flag = ?", false)
		databaseHelper.
			On("Where", mock.AnythingOfType("string"), mock.AnythingOfType("bool")).
			Return(databaseHelper)
		// Order("id DESC")
		databaseHelper.
			On("Order", mock.AnythingOfType("string")).
			Return(databaseHelper)
		// Find(&photos)
		databaseHelper.
			On("Find", mock.AnythingOfType("*[]domain.Photo")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*[]domain.Photo)
				*out = *photosPtr
			}).
			Return(nil)

		pr := repository.NewPhotoRepository(databaseHelper)

		photos, err := pr.FindAllAvailable()

		// Kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, expectedPhotos, photos)
		assert.Len(t, photos, len(expectedPhotos))

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một mock cho database
		databaseHelper := &mocks.Database{}

		expectedPhotos := []domain.Photo{}

		photosPtr := &expectedPhotos

		// behavior
		// p.db.Model(&domain.Photo{}).Where("delete_flag = ?", false).Order("id DESC").Find(&photos)

		// Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper)
		// Where("delete_flag = ?", false)
		databaseHelper.
			On("Where", mock.AnythingOfType("string"), mock.AnythingOfType("bool")).
			Return(databaseHelper)
		// Order("id DESC")
		databaseHelper.
			On("Order", mock.AnythingOfType("string")).
			Return(databaseHelper)
		// Find(&photos)
		databaseHelper.
			On("Find", mock.AnythingOfType("*[]domain.Photo")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*[]domain.Photo)
				*out = *photosPtr
			}).
			Return(gorm.ErrRecordNotFound)

		pr := repository.NewPhotoRepository(databaseHelper)
		photos, err := pr.FindAllAvailable()

		// Kiểm tra kết quả
		assert.Error(t, err)
		assert.Len(t, photos, len(expectedPhotos))
		assert.Equal(t, expectedPhotos, photos)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})
}

func TestGetAllAvailableID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Tạo một mock cho database
		databasehelper := &mocks.Database{}

		expectedIDs := []string{"1", "2"}

		idsPtr := &expectedIDs

		// behavior
		//p.db.Model(&domain.Photo{}).Select("id").Where("delete_flag = ?", false).Order("id DESC").Find(&ids)

		// Model(&domain.Photo{})
		databasehelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databasehelper)
		// Select("id")
		databasehelper.
			On("Select", mock.AnythingOfType("string")).
			Return(databasehelper)
		// Where("delete_flag = ?", false)
		databasehelper.
			On("Where", mock.AnythingOfType("string"), mock.AnythingOfType("bool")).
			Return(databasehelper)
		// Order("id DESC")
		databasehelper.
			On("Order", mock.AnythingOfType("string")).
			Return(databasehelper)
		// Find(&ids)
		databasehelper.
			On("Find", mock.AnythingOfType("*[]string")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*[]string)
				*out = *idsPtr
			}).
			Return(nil)

		pr := repository.NewPhotoRepository(databasehelper)
		ids, err := pr.GetAllAvailableID()

		// Kiểm tra kết quả
		assert.NoError(t, err)
		assert.Len(t, ids, len(expectedIDs))
		assert.Equal(t, expectedIDs, ids)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databasehelper.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một mock cho database
		databasehelper := &mocks.Database{}

		expectedIDs := []string{}

		idsPtr := &expectedIDs

		// behavior
		//p.db.Model(&domain.Photo{}).Select("id").Where("delete_flag = ?", false).Order("id DESC").Find(&ids)

		// Model(&domain.Photo{})
		databasehelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databasehelper)
		// Select("id")
		databasehelper.
			On("Select", mock.AnythingOfType("string")).
			Return(databasehelper)
		// Where("delete_flag = ?", false)
		databasehelper.
			On("Where", mock.AnythingOfType("string"), mock.AnythingOfType("bool")).
			Return(databasehelper)
		// Order("id DESC")
		databasehelper.
			On("Order", mock.AnythingOfType("string")).
			Return(databasehelper)
		// Find(&ids)
		databasehelper.
			On("Find", mock.AnythingOfType("*[]string")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*[]string)
				*out = *idsPtr
			}).
			Return(gorm.ErrRecordNotFound)

		pr := repository.NewPhotoRepository(databasehelper)
		ids, err := pr.GetAllAvailableID()

		// Kiểm tra kết quả
		assert.Error(t, err)
		assert.Len(t, ids, len(expectedIDs))
		assert.Equal(t, expectedIDs, ids)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databasehelper.AssertExpectations(t)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Tạo một mock cho database
		databaseHelper := &mocks.Database{}

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
				DeleteFlag:  newFalse(),
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		}

		photosPtr := &expectedPhotos

		// behavior
		// p.db.Model(&domain.Photo{}).Order("id DESC").Find(&photos)

		// Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper)
		// Order("id DESC")
		databaseHelper.
			On("Order", mock.AnythingOfType("string")).
			Return(databaseHelper)
		// Find(&photos)
		databaseHelper.
			On("Find", mock.AnythingOfType("*[]domain.Photo")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*[]domain.Photo)
				*out = *photosPtr
			}).
			Return(nil)

		pr := repository.NewPhotoRepository(databaseHelper)
		photos, err := pr.FindAll()

		// Kiểm tra kết quả
		assert.NoError(t, err)
		assert.Len(t, photos, len(expectedPhotos))
		assert.Equal(t, expectedPhotos, photos)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một mock cho database
		databaseHelper := &mocks.Database{}

		expectedPhotos := []domain.Photo{}

		photosPtr := &expectedPhotos

		// behavior
		//p.db.Model(&domain.Photo{}).Order("id DESC").Find(&photos)

		// Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper)
		// Order("id DESC")
		databaseHelper.
			On("Order", mock.AnythingOfType("string")).
			Return(databaseHelper)
		// Find(&photos)
		databaseHelper.
			On("Find", mock.AnythingOfType("*[]domain.Photo")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*[]domain.Photo)
				*out = *photosPtr
			}).
			Return(gorm.ErrRecordNotFound)

		pr := repository.NewPhotoRepository(databaseHelper)
		photos, err := pr.FindAll()

		// Kiểm tra kết quả
		assert.Error(t, err)
		assert.Len(t, photos, len(expectedPhotos))
		assert.Equal(t, expectedPhotos, photos)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})
}

func TestUpdateByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Tạo một mock cho database
		databaseHelper := &mocks.Database{}

		photo := domain.Photo{
			ID:          uint(1),
			Name:        "test update",
			Url:         "https://test.png",
			Description: "test",
			DeleteFlag:  newTrue(),
			CreatedAt:   now,
			UpdatedAt:   time.Now(),
		}

		photoPtr := &photo

		// behavior
		// p.db.Model(&domain.Photo{}).Where("id = ?", photo.ID).Updates(photo)

		// Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper)
		// Where("id = ?", photo.ID)
		databaseHelper.
			On("Where", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).
			Return(databaseHelper)
		// Updates(photo)
		databaseHelper.
			On("Updates", mock.AnythingOfType("*domain.Photo")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*domain.Photo)
				*out = *photoPtr
			}).
			Return(nil)

		pr := repository.NewPhotoRepository(databaseHelper)
		updatedPhoto, err := pr.UpdateByID(photo)

		// kiểm tra kết quả
		assert.NoError(t, err)
		assert.Equal(t, photo, updatedPhoto)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một mock cho database
		databaseHelper := &mocks.Database{}

		photo := domain.Photo{
			ID:          uint(2),
			Name:        "test update",
			Url:         "https://test.png",
			Description: "test",
			DeleteFlag:  newTrue(),
			CreatedAt:   now,
			UpdatedAt:   time.Now(),
		}

		photoPtr := &photo

		// behavior
		// p.db.Model(&domain.Photo{}).Where("id = ?", photo.ID).Updates(&photo)

		// Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper)
		// Where("id = ?", photo.ID)
		databaseHelper.
			On("Where", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).
			Return(databaseHelper)
		// Updates(&photo)
		databaseHelper.
			On("Updates", mock.AnythingOfType("*domain.Photo")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*domain.Photo)
				*out = *photoPtr
			}).
			Return(gorm.ErrRecordNotFound)

		pr := repository.NewPhotoRepository(databaseHelper)
		updatedPhoto, err := pr.UpdateByID(photo)

		// kiểm tra kết quả
		assert.Error(t, err)
		assert.Equal(t, photo.Name, updatedPhoto.Name)
		assert.Equal(t, photo.Url, updatedPhoto.Url)
		assert.Equal(t, photo.Description, updatedPhoto.Description)
		assert.Equal(t, photo.DeleteFlag, updatedPhoto.DeleteFlag)
		assert.Equal(t, photo.CreatedAt, updatedPhoto.CreatedAt)
		assert.Equal(t, photo.UpdatedAt, updatedPhoto.UpdatedAt)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})
}

func TestInsertPhoto(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Tạo một mock cho database
		databaseHelper := &mocks.Database{}

		photo := domain.Photo{
			Name:        "test update",
			Url:         "https://test.png",
			Description: "test",
		}

		insertedPhoto := domain.Photo{
			ID:          uint(1),
			Name:        "test update",
			Url:         "https://test.png",
			Description: "test",
			DeleteFlag:  newFalse(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		photoPtr := &insertedPhoto

		//behavior
		//  p.db.Model(&domain.Photo{}).Create(&photo)

		// Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper)
		// Create(&photo)
		databaseHelper.
			On("Create", mock.AnythingOfType("*domain.Photo")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*domain.Photo)
				*out = *photoPtr
			}).
			Return(nil)

		pr := repository.NewPhotoRepository(databaseHelper)
		resultPhoto, err := pr.InsertPhoto(photo)

		// kiểm tra kiêt1 quả
		assert.NoError(t, err)
		assert.Equal(t, resultPhoto.Name, insertedPhoto.Name)
		assert.Equal(t, resultPhoto.Url, insertedPhoto.Url)
		assert.Equal(t, resultPhoto.Description, insertedPhoto.Description)
		assert.Equal(t, resultPhoto.DeleteFlag, insertedPhoto.DeleteFlag)
		assert.Equal(t, resultPhoto.CreatedAt, insertedPhoto.CreatedAt)
		assert.Equal(t, resultPhoto.UpdatedAt, insertedPhoto.UpdatedAt)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		databaseHelper := &mocks.Database{}

		photo := domain.Photo{
			Name:        "test update",
			Url:         "",
			Description: "test",
		}

		photoPtr := &photo

		expectedError := errors.New("null value in column \"url\" violates not-null constraint")

		//behavior
		//  p.db.Model(&domain.Photo{}).Create(&photo)

		// Model(&domain.Photo{})
		databaseHelper.
			On("Model", mock.AnythingOfType("*domain.Photo")).
			Return(databaseHelper)
		// Create(&photo)
		databaseHelper.
			On("Create", mock.AnythingOfType("*domain.Photo")).
			Run(func(args mock.Arguments) {
				out := args.Get(0).(*domain.Photo)
				*out = *photoPtr
			}).
			Return(expectedError)

		pr := repository.NewPhotoRepository(databaseHelper)
		resultPhoto, err := pr.InsertPhoto(photo)

		// kiểm tra kết quả

		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedError.Error())

		assert.Equal(t, resultPhoto.Name, photo.Name)
		assert.Equal(t, resultPhoto.Url, photo.Url)
		assert.Equal(t, resultPhoto.Description, photo.Description)

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		databaseHelper.AssertExpectations(t)
	})
}
