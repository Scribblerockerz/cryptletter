package web

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
)

//go:embed dist/*
var Assets embed.FS

// fsFunc is short-hand for constructing a http.FileSystem
type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

// AssetHandler returns an http.Handler that will serve files from Assets embed.FS
func AssetHandler(prefix, root string) http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)

		f, err := Assets.Open(assetPath)
		if os.IsNotExist(err) {
			return Assets.Open("dist/index.html")
		}

		return f, err
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}
