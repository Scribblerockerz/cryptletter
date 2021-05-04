package web

import (
	"bytes"
	"embed"
	"github.com/spf13/viper"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
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
			f, err = Assets.Open("dist/index.html")
		}

		f, err = injectAdditionalContent(f)

		return f, err
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}


// injectAdditionalContent will modify file contents, and inject additional assets by configuration
func injectAdditionalContent(file fs.File) (fs.File, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileInfo.Name() != "index.html" {
		return file, nil
	}

	contentBytes := make([]byte, fileInfo.Size())
	n1, err := file.Read(contentBytes)

	var additionalContent string

	// Inject CSS
	cssPath := viper.GetString("app.additional.css")
	if cssPath != "" {
		content, err := ioutil.ReadFile(cssPath)
		if err == nil {
			additionalContent += "<style>" + string(content) + "</style>"
		}
	}

	// Inject JS
	jsPath := viper.GetString("app.additional.js")
	if jsPath != "" {
		content, err := ioutil.ReadFile(jsPath)
		if err == nil {
			additionalContent += "<script>" + string(content) + "</script>"
		}
	}

	content := string(contentBytes[:n1])

	if additionalContent != "" {
		content = strings.Replace(string(contentBytes[:n1]), "</head>", additionalContent + "</head>", -1)
	}

	patchedFile := inMemoryFile{
		inMemoryFileInfo: inMemoryFileInfo{fileInfo: fileInfo, size: int64(len(content))},
		buf: bytes.NewBufferString(content),
	}

	return patchedFile, nil
}

type inMemoryFile struct {
	fs.File
	inMemoryFileInfo inMemoryFileInfo // Size() will return an inaccurate value, since we modified it
	buf *bytes.Buffer
}

func (im inMemoryFile) Stat() (fs.FileInfo, error) {
	return im.inMemoryFileInfo, nil
}

func (im inMemoryFile) Read(b []byte) (int, error) {
	return im.buf.Read(b)
}

func (im inMemoryFile) Close() error {
	return nil
}

type inMemoryFileInfo struct {
	fs.FileInfo
	fileInfo fs.FileInfo
	size int64
}

func (fi inMemoryFileInfo) Name() string {
	return fi.fileInfo.Name()
}

func (fi inMemoryFileInfo) Size() int64 {
	return fi.size
}

func (fi inMemoryFileInfo) Mode() fs.FileMode {
	return fi.fileInfo.Mode()
}

func (fi inMemoryFileInfo) ModTime() time.Time {
	return fi.fileInfo.ModTime()
}

func (fi inMemoryFileInfo) IsDir() bool {
	return fi.fileInfo.IsDir()
}
