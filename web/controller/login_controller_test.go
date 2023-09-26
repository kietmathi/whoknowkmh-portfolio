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
	"github.com/kietmathi/whoknowkmh-portfolio/bootstrap"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/domain/mocks"
	"github.com/kietmathi/whoknowkmh-portfolio/web/controller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginController_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		// mock usecase & controller
		mockLoginUsecase := new(mocks.LoginUsecase)
		loginController := &controller.LoginController{
			LoginUsecase: mockLoginUsecase,
		}

		// behavior
		mockLoginUsecase.On("GetSession", mock.AnythingOfType("*gin.Context"), "error").Return(nil, nil)
		mockLoginUsecase.On("DeleteFromSession", mock.AnythingOfType("*gin.Context"), "error").Return(nil)
		mockLoginUsecase.On(
			"RenderTemplate",
			mock.AnythingOfType("*gin.Context"), // Kiểm tra kiểu của tham số đầu tiên
			http.StatusOK,
			domain.LoginTemplateName,
			0*time.Second, // Kiểm tra cacheDuration
			mock.Anything, // Kiểm tra dữ liệu (data) là bất kỳ giá trị nào
		).Return()

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.GET("/login", loginController.Login)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("GET", "/login", nil)
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusOK, w.Code)
		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		mockLoginUsecase.AssertExpectations(t)
	})
}

func TestLoginController_LoginPost(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		// mock usecase, env & controller
		mockLoginUsecase := new(mocks.LoginUsecase)
		mockEnv := &bootstrap.Env{
			AdminUserID:      "admin",
			AdminUserPwdHash: "$2a$12$IPeojFPGG5s3kKNGxWgWouuibUWRo9OJUrs3nZIbp.yZnGkSvuLyS",
			// Các giá trị khác của Env
		}
		loginController := &controller.LoginController{
			LoginUsecase: mockLoginUsecase,
			Env:          mockEnv,
		}

		// mock login user
		mockLoginUser := domain.LoginUser{
			UserID:   "admin",
			Password: "hashed_password",
		}

		// behavior
		mockLoginUsecase.On("CreateAccessToken", mock.AnythingOfType("*domain.LoginUser"), mock.Anything, mock.Anything).Return("access_token", nil)
		mockLoginUsecase.On("DeleteFromSession", mock.AnythingOfType("*gin.Context"), "error").Return(nil)
		mockLoginUsecase.On("SetCookieSession", mock.AnythingOfType("*gin.Context"), "Authorization", "access_token", mock.Anything).Return(nil)

		// request login user
		requestLoginUser, err := json.Marshal(mockLoginUser)
		assert.NoError(t, err)

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.POST("/login", loginController.LoginPost)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestLoginUser))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusSeeOther, w.Code)
		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		assert.Equal(t, "/admin", w.Header().Get("Location"))
	})

	t.Run("errorBind", func(t *testing.T) {

		// mock user, logger & controller
		mockLoginUsecase := new(mocks.LoginUsecase)
		mockLogger := new(mocks.Logger)
		loginController := &controller.LoginController{
			LoginUsecase: mockLoginUsecase,
			Logger:       mockLogger,
		}

		// mock login user
		mockLoginUser := domain.LoginUser{
			UserID:   "admin",
			Password: "",
		}

		// behavior
		expectedErr := errors.New("Key: 'LoginUser.Password' Error:Field validation for 'Password' failed on the 'required' tag")
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()
		mockLoginUsecase.On("SetSession", mock.AnythingOfType("*gin.Context"), "error", expectedErr.Error()).Return(nil)

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.POST("/login", loginController.LoginPost)

		// request login user
		requestLoginUser, err := json.Marshal(mockLoginUser)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestLoginUser))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusSeeOther, w.Code)
		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		assert.Equal(t, "/login", w.Header().Get("Location"))
	})

	t.Run("errorFailedCredential", func(t *testing.T) {

		// mock usecase, logger, env & controller
		mockLoginUsecase := new(mocks.LoginUsecase)
		mockLogger := new(mocks.Logger)
		mockEnv := &bootstrap.Env{
			AdminUserID:      "admin",
			AdminUserPwdHash: "$2a$12$IPeojFPGG5s3kKNGxWgWouuibUWRo9OJUrs3nZIbp.yZnGkSvuLyS", // hashed_password
		}
		loginController := &controller.LoginController{
			LoginUsecase: mockLoginUsecase,
			Logger:       mockLogger,
			Env:          mockEnv,
		}

		// mock login user
		mockLoginUser := domain.LoginUser{
			UserID:   "admin",
			Password: "incorrect password",
		}

		// behavior
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()
		mockLoginUsecase.On("SetSession", mock.AnythingOfType("*gin.Context"), "error", "Invalid login credentials").Return(nil)

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.POST("/login", loginController.LoginPost)

		// request login user
		requestLoginUser, err := json.Marshal(mockLoginUser)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestLoginUser))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusSeeOther, w.Code)
		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		assert.Equal(t, "/login", w.Header().Get("Location"))
	})

	t.Run("errorCreateAccessToken", func(t *testing.T) {

		// mock usecase, logger, env & controller
		mockLoginUsecase := new(mocks.LoginUsecase)
		mockLogger := new(mocks.Logger)
		mockEnv := &bootstrap.Env{
			AdminUserID:      "admin",
			AdminUserPwdHash: "$2a$12$IPeojFPGG5s3kKNGxWgWouuibUWRo9OJUrs3nZIbp.yZnGkSvuLyS",
		}
		loginController := &controller.LoginController{
			LoginUsecase: mockLoginUsecase,
			Logger:       mockLogger,
			Env:          mockEnv,
		}

		// mock login user
		mockLoginUser := domain.LoginUser{
			UserID:   "admin",
			Password: "hashed_password",
		}

		// behavior
		mockError := errors.New("test error")
		mockLoginUsecase.On("CreateAccessToken", mock.AnythingOfType("*domain.LoginUser"), mock.Anything, mock.Anything).Return("", mockError)
		mockLogger.On("Printf", mock.AnythingOfType("string"), mock.Anything).Return()
		mockLoginUsecase.On("SetSession", mock.AnythingOfType("*gin.Context"), "error", mockError.Error()).Return(nil)

		// tạo gin router
		router := gin.Default()
		// Định nghĩa endpoint API và gọi controller method
		router.POST("/login", loginController.LoginPost)

		// request login user
		requestLoginUser, err := json.Marshal(mockLoginUser)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		// gọi endpoint và kiểm tra response
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestLoginUser))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusSeeOther, w.Code)
		// Đảm bảo rằng phương thức đã được gọi theo mong muốn
		assert.Equal(t, "/login", w.Header().Get("Location"))
	})
}
