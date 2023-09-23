package usecase_test

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/stretchr/testify/assert"
)

func TestLoginUsecaseSetSessionWithGin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Tạo một bộ kiểm tra HTTP và router Gin
		router := gin.Default()
		// set store
		store := cookie.NewStore([]byte("secret"))
		router.Use(sessions.Sessions("mysession", store))

		// Định nghĩa một endpoint API để kiểm tra use case
		router.GET("/set-session", func(c *gin.Context) {
			// Tạo use case và gọi hàm SetSession
			loginUsecase := usecase.NewLoginUsecase()
			key := "example"
			value := "data"
			err := loginUsecase.SetSession(c, key, value)

			// Kiểm tra lỗi
			assert.NoError(t, err)

			// Kiểm tra xem session đã được đặt chính xác hay không
			session := sessions.Default(c)
			sessionValue := session.Get(key)
			assert.Equal(t, value, sessionValue)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/set-session", nil)
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một bộ kiểm tra HTTP và router Gin
		router := gin.Default()

		// Định nghĩa một điểm cuối API để kiểm tra use case
		router.GET("/set-session", func(c *gin.Context) {
			// Tạo use case và gọi hàm SetSession khi chưa set store
			loginUsecase := usecase.NewLoginUsecase()
			key := "example"
			value := "data"
			err := loginUsecase.SetSession(c, key, value)

			// Kiểm tra lỗi
			assert.Error(t, err)

		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/set-session", nil)
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestLoginUsecaseGetSessionWithGin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Tạo một bộ kiểm tra HTTP và router Gin
		router := gin.Default()
		// set store
		store := cookie.NewStore([]byte("secret"))
		router.Use(sessions.Sessions("mysession", store))

		// Định nghĩa một endpoint API để kiểm tra use case
		router.GET("/get-session", func(c *gin.Context) {

			// set session
			key := "example"
			value := "data"
			session := sessions.Default(c)
			session.Set(key, value)
			session.Save()

			// Tạo use case và gọi hàm GetSession
			loginUsecase := usecase.NewLoginUsecase()
			sessionValue, err := loginUsecase.GetSession(c, key)

			// Kiểm tra lỗi và session get được chính xác hay không
			assert.NoError(t, err)
			assert.Equal(t, value, sessionValue)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/get-session", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một bộ kiểm tra HTTP và router Gin
		router := gin.Default()

		// Định nghĩa một endpoint API để kiểm tra use case
		router.GET("/get-session", func(c *gin.Context) {

			// set session khi store chưa được set
			key := "example"
			value := "data"
			session := sessions.Default(c)
			session.Set(key, value)
			session.Save()

			// Tạo use case và gọi hàm GetSession khi store chưa được set
			loginUsecase := usecase.NewLoginUsecase()
			_, err := loginUsecase.GetSession(c, key)

			// Kiểm tra lỗi
			assert.Error(t, err)

		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/get-session", nil)
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestLoginUsecaseDeleteFromSessionWithGin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Tạo một bộ kiểm tra HTTP và router Gin
		router := gin.Default()
		// set store
		store := cookie.NewStore([]byte("secret"))
		router.Use(sessions.Sessions("mysession", store))

		// Định nghĩa một endpoint API để kiểm tra use case
		router.GET("/get-session", func(c *gin.Context) {

			// set session
			key := "example"
			value := "data"
			session := sessions.Default(c)
			session.Set(key, value)
			session.Save()

			// Tạo use case và gọi hàm DeleteFromSession
			loginUsecase := usecase.NewLoginUsecase()
			err := loginUsecase.DeleteFromSession(c, key)

			// Kiểm tra lỗi
			assert.NoError(t, err)

			// get session
			sessionValue := session.Get(key)
			// Kiểm tra xem session đã được get chính xác hay không
			assert.Nil(t, sessionValue)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/get-session", nil)
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một bộ kiểm tra HTTP và router Gin và không set session store
		router := gin.Default()

		// Định nghĩa một endpoint API để kiểm tra use case
		router.GET("/delete-session", func(c *gin.Context) {

			// set session khi store chưa được set
			key := "example"
			value := "data"
			session := sessions.Default(c)
			session.Set(key, value)
			session.Save()

			// Tạo use case và gọi hàm DeleteFromSession khi store chưa được set
			loginUsecase := usecase.NewLoginUsecase()
			err := loginUsecase.DeleteFromSession(c, key)

			// Kiểm tra lỗi
			assert.Error(t, err)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/delete-session", nil)
		router.ServeHTTP(w, req)

		// kiểm tra status code
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestLoginUsecaseSetCookieSessionWithGin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Tạo một bộ kiểm tra HTTP và router Gin
		router := gin.Default()

		// thông tin cockie
		key := "example"
		value := "data"
		maxAge := 3600

		// Định nghĩa một endpoint API để kiểm tra use case
		router.GET("/set-cockie", func(c *gin.Context) {

			// Tạo use case và gọi hàm SetCookieSession
			loginUsecase := usecase.NewLoginUsecase()
			err := loginUsecase.SetCookieSession(c, key, value, maxAge)

			// Kiểm tra lỗi
			assert.NoError(t, err)

		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/set-cockie", nil)
		router.ServeHTTP(w, req)

		// kiểm tra thông tin cockie
		cookie := w.Result().Cookies()[0]
		assert.Equal(t, key, cookie.Name)
		assert.Equal(t, value, cookie.Value)
		assert.Equal(t, "/", cookie.Path)
		assert.Equal(t, maxAge, cookie.MaxAge)

		// kiểm tra status code
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestLoginCreateAccessToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := &domain.LoginUser{
			UserID:   "admin",
			Password: "password",
		}

		secret := "secret_key"
		expiry := 1

		loginUsecase := usecase.NewLoginUsecase()
		accessToken, err := loginUsecase.CreateAccessToken(user, secret, expiry)

		// Kiểm tra lỗi, nếu không có lỗi thì lỗi phải là nil
		assert.NoError(t, err)

		// Kiểm tra xem accessToken có được tạo ra không
		assert.NotEmpty(t, accessToken)

		// Kiểm tra xem accessToken có đúng định dạng JWT không
		token, parseErr := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
		assert.NoError(t, parseErr)
		assert.True(t, token.Valid)

		// Kiểm tra xem accessToken có chứa claims (id và exp) đúng không
		claims, ok := token.Claims.(jwt.MapClaims)
		assert.True(t, ok)
		assert.Equal(t, user.UserID, claims["id"])
		expTime, ok := claims["exp"].(float64)
		assert.True(t, ok)

		// Kiểm tra thời gian hết hạn trong accessToken

		expUnix := int64(expTime)
		expectedExpirationTime := time.Now().Add(time.Hour * time.Duration(expiry))
		expectedExpUnix := expectedExpirationTime.Unix()
		assert.Equal(t, expUnix, expectedExpUnix)
	})
}

func TestLoginUsecaseRenderTemplateWithGin(t *testing.T) {
	tpl := "template.html"

	t.Run("success", func(t *testing.T) {
		// Tạo một đối tượng use case
		lu := usecase.NewLoginUsecase()

		// Tạo một bộ kiểm tra HTTP và router Gin
		r := gin.Default()

		// set template
		templ := template.Must(template.New(tpl).Parse(`Hello {{.data.title}}`))
		r.SetHTMLTemplate(templ)

		// Định nghĩa một endpoint API để kiểm tra use case
		r.GET("/render-template", func(c *gin.Context) {
			// render template
			data := make(map[string]interface{}, 1)
			data["title"] = domain.LoginTitle
			lu.RenderTemplate(
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
		assert.Equal(t, "Hello "+domain.LoginTitle, w.Body.String())
		assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một đối tượng use case
		u := usecase.NewLoginUsecase()

		// Tạo một bộ kiểm tra HTTP và router Gin
		r := gin.Default()

		// Định nghĩa một endpoint API để kiểm tra use case
		r.GET("/render-template", func(c *gin.Context) {
			// render với template không tồn tại (chưa được set)
			data := make(map[string]interface{}, 1)
			data["title"] = domain.LoginTitle
			u.RenderTemplate(
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
