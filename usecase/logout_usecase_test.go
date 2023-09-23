package usecase_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/stretchr/testify/assert"
)

func TestLogoutUsecaseSetSessionWithGin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := gin.Default()
		store := cookie.NewStore([]byte("secret"))
		router.Use(sessions.Sessions("mysession", store))

		// Định nghĩa một điểm cuối API để kiểm tra use case
		router.GET("/set-session", func(c *gin.Context) {
			// Giả lập việc middleware đã xử lý CSRF và đã đặt giá trị cho csrf.TemplateTag
			// Tạo use case và gọi hàm SetSession
			logoutUsecase := usecase.NewLogoutUsecase() // Điều này cần phải được thay thế bằng cách khởi tạo use case thực tế của bạn.
			key := "example"
			value := "data"
			err := logoutUsecase.SetSession(c, key, value)

			// Kiểm tra lỗi và session đã được đặt chính xác hay không
			assert.NoError(t, err)

			// Kiểm tra xem session đã được đặt chính xác hay không
			session := sessions.Default(c)
			sessionValue := session.Get(key)
			assert.Equal(t, value, sessionValue)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/set-session", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error", func(t *testing.T) {
		router := gin.Default()
		// store := cookie.NewStore([]byte("secret"))
		// router.Use(sessions.Sessions("mysession", store))

		// Định nghĩa một điểm cuối API để kiểm tra use case
		router.GET("/set-session", func(c *gin.Context) {
			// Giả lập việc middleware đã xử lý CSRF và đã đặt giá trị cho csrf.TemplateTag
			// Tạo use case và gọi hàm SetSession
			logoutUsecase := usecase.NewLogoutUsecase() // Điều này cần phải được thay thế bằng cách khởi tạo use case thực tế của bạn.
			key := "example"
			value := "data"
			err := logoutUsecase.SetSession(c, key, value)

			// Kiểm tra lỗi và session đã được đặt chính xác hay không
			assert.Error(t, err)

		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/set-session", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestLogoutUsecaseDeleteFromCookieSessionWithGin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		router := gin.Default()

		key := "example"
		value := "data"
		maxAge := 3600

		// Định nghĩa một điểm cuối API để kiểm tra use case
		router.GET("/delete-cockie", func(c *gin.Context) {

			// Tạo use case và gọi hàm SetSession
			logoutUsecase := usecase.NewLogoutUsecase() // Điều này cần phải được thay thế bằng cách khởi tạo use case thực tế của bạn.

			err := logoutUsecase.DeleteFromCookieSession(c, key)

			// Kiểm tra lỗi và session đã được đặt chính xác hay không
			assert.NoError(t, err)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/delete-cockie", nil)
		reqCookie := http.Cookie{
			Name:     key,
			Value:    value,
			MaxAge:   maxAge,
			SameSite: http.SameSiteLaxMode,
		}
		req.AddCookie(&reqCookie)
		router.ServeHTTP(w, req)

		rspCookie := w.Result().Cookies()[0]
		assert.Equal(t, key, rspCookie.Name)
		assert.Equal(t, "", rspCookie.Value)
		assert.Equal(t, "/", rspCookie.Path)
		assert.Equal(t, -1, rspCookie.MaxAge)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
