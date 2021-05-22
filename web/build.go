package web

import (
	"bytes"
	"embed"
	"encoding/json"
	"github.com/Scribblerockerz/cryptletter/pkg/utils"
	"github.com/spf13/viper"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
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

// frontendOptions are used to provide configuration to the frontend
type frontendOptions struct {
	SupportsAttachments bool `json:"supportsAttachments"`
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
	additionalContent += getInjectableCSS()

	// Inject JS Options
	additionalContent += getInjectableOptions()

	// Inject JS
	additionalContent += getInjectableJS()

	content := string(contentBytes[:n1])

	if additionalContent != "" {
		content = strings.Replace(string(contentBytes[:n1]), "</head>", additionalContent+"</head>", -1)
	}

	patchedFile := utils.InMemoryFile{
		InMemoryFileInfo: utils.InMemoryFileInfo{FileInfoRef: fileInfo, FileSize: int64(len(content))},
		Buf:              bytes.NewBufferString(content),
	}

	return patchedFile, nil
}

func getInjectableCSS() string {
	cssPath := viper.GetString("app.additional.css")
	if cssPath != "" {
		content, err := ioutil.ReadFile(cssPath)
		if err == nil {
			return "<style>" + string(content) + "</style>"
		}
	}

	return ""
}

func getInjectableOptions() string {
	jsOptions, err := json.Marshal(frontendOptions{
		SupportsAttachments: viper.GetString("app.attachments.driver") != "",
	})
	if err != nil {
		return ""
	}
	return "<script>window.cryptletterOptions = " + string(jsOptions) + "</script>"
}

func getInjectableJS() string {
	jsPath := viper.GetString("app.additional.js")
	if jsPath != "" {
		content, err := ioutil.ReadFile(jsPath)
		if err == nil {
			return "<script>" + string(content) + "</script>"
		}
	}

	return ""
}
