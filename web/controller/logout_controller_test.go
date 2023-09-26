package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain/mocks"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogoutController_Logout(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		// mock usecase, logger & controller
		mockLogoutUsecase := new(mocks.LogoutUsecase)
		mockLogger := new(mocks.Logger)
		logoutController := &controller.LogoutController{
			LogoutUsecase: mockLogoutUsecase,
			Logger:        mockLogger,
		}

		// behavior
		mockLogoutUsecase.On("DeleteFromCookieSession", mock.AnythingOfType("*gin.Context"), "Authorization").Return(nil)

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/logout", logoutController.Logout)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/logout", nil)
		router.ServeHTTP(w, req)

		// Kiểm tra kết quả
		assert.Equal(t, http.StatusSeeOther, w.Code)
		assert.Equal(t, "/login", w.Header().Get("Location"))

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockLogoutUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {

		// mock usecase, logger & controller
		mockLogoutUsecase := new(mocks.LogoutUsecase)
		mockLogger := new(mocks.Logger)
		logoutController := &controller.LogoutController{
			LogoutUsecase: mockLogoutUsecase,
			Logger:        mockLogger,
		}

		// behavior
		expectedError := errors.New("test error")
		mockLogoutUsecase.On("DeleteFromCookieSession", mock.AnythingOfType("*gin.Context"), "Authorization").Return(errors.New("test error"))
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()
		mockLogoutUsecase.On("SetSession", mock.AnythingOfType("*gin.Context"), "error", expectedError.Error()).Return(nil)

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/logout", logoutController.Logout)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/logout", nil)
		router.ServeHTTP(w, req)

		// Kiểm tra kết quả
		assert.Equal(t, http.StatusSeeOther, w.Code)
		assert.Equal(t, "/login", w.Header().Get("Location"))

		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockLogoutUsecase.AssertExpectations(t)
	})

}
