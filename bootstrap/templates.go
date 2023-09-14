package bootstrap

import (
	"embed"
	"html/template"
	"log"

	"github.com/gin-contrib/multitemplate"
)

// NewEmbedTemplates : Make a multitemplate.Renderer instance from
// the template stored in our embedded file system to render dynamic content
func NewEmbedTemplates(fs embed.FS) multitemplate.Renderer {
	renderer := multitemplate.New()

	// Generate our templates map from our layouts/ and includes/ directories

	layouts, err := embed.FS.ReadDir(fs, "templates/layouts")
	if err != nil {
		log.Fatal(err)
	}

	for _, layout := range layouts {
		embeddedTemplate, err :=
			template.ParseFS(fs, "templates/includes/base.html", "templates/layouts/"+layout.Name())
		if err != nil {
			log.Fatal(err)
		}
		renderer.Add(layout.Name(), embeddedTemplate)
		log.Println(layout.Name() + " loaded")
	}

	return renderer
}
