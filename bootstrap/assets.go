package bootstrap

import (
	"embed"
	"io/fs"
	"log"
)

func NewEmbedAssets(embedFS embed.FS) fs.FS {
	assets, err := fs.Sub(embedFS, "assets")
	if err != nil {
		log.Fatal(err)
	}
	return assets
}
