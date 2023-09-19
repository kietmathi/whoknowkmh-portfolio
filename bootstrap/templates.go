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

	dirs, err := embed.FS.ReadDir(fs, "templates")
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		layouts, err := embed.FS.ReadDir(fs, "templates/"+dir.Name()+"/layouts")
		if err != nil {
			log.Fatal(err)
		}

		for _, layout := range layouts {
			embeddedTemplate, err :=
				template.ParseFS(fs, "templates/"+dir.Name()+"/includes/base.html", "templates/"+dir.Name()+"/layouts/"+layout.Name())
			if err != nil {
				log.Fatal(err)
			}
			renderer.Add(dir.Name()+"/"+layout.Name(), embeddedTemplate)
			log.Println(dir.Name() + "/" + layout.Name() + " loaded")
		}
	}

	return renderer
}
