package usecase_test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kietmathi/whoknowkmh-portfolio/domain"
	"github.com/kietmathi/whoknowkmh-portfolio/usecase"
	"github.com/stretchr/testify/assert"
)

func TestBlogUsecaseRenderTemplateWithGin(t *testing.T) {
	tpl := "template.html"

	t.Run("success", func(t *testing.T) {
		// Tạo một đối tượng use case mới
		u := usecase.NewBlogUsecase()

		// Tạo một bộ kiểm tra HTTP và router Gin
		r := gin.Default()

		// set template
		templ := template.Must(template.New(tpl).Parse(`Hello {{.data.title}}`))
		r.SetHTMLTemplate(templ)

		// Định nghĩa một endpoint API để kiểm tra use case
		r.GET("/render-template", func(c *gin.Context) {
			// render template
			data := make(map[string]interface{}, 1)
			data["title"] = domain.BlogTitle
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
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Hello "+domain.BlogTitle, w.Body.String())
		assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	})

	t.Run("error", func(t *testing.T) {
		// Tạo một đối tượng use case
		u := usecase.NewBlogUsecase()

		// Tạo một bộ kiểm tra HTTP và router Gin
		r := gin.Default()

		// Định nghĩa một endpoint API để kiểm tra use case
		r.GET("/render-template", func(c *gin.Context) {
			// render với template không tồn tại (chưa được set)
			data := make(map[string]interface{}, 1)
			data["title"] = domain.BlogTitle
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

		// Kiểm tra xem phản ứng có đúng không
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
