package template

import (
	"fmt"
	"github.com/Scribblerockerz/cryptletter/pkg/logger"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/aymerick/raymond"
)

// RenderLayout renders a layout path with a context
func RenderLayout(path string, ctx map[string]string) string {
	tpl, err := raymond.ParseFile(resolveTemplatePath(path))
	if err != nil {
		panic(err)
	}

	result := tpl.MustExec(ctx)
	return result
}

func resolveTemplatePath(path string) string {
	templateDir := "./web/theme/templates" //main.Config.App.TemplatesDir
	fallbackDir := "./web/theme/templates" //main.DefaultConfiguration.App.TemplatesDir

	if _, err := os.Stat(templateDir + "/" + path); err == nil {
		return templateDir + "/" + path

	} else if os.IsNotExist(err) {
		return fallbackDir + "/" + path
	} else {
		panic(err)
	}
}

// RegisterPartials will scan the template dir and the provided over-ride dir
func RegisterPartials() {
	templateDir := "./web/theme/templates" //main.Config.App.TemplatesDir
	fallbackDir := "./web/theme/templates" //main.DefaultConfiguration.App.TemplatesDir

	var partialList []string

	if templateDir != fallbackDir {
		templateDirPartials := scanPartialsInPath(templateDir, "")
		fallbackDirPartials := scanPartialsInPath(fallbackDir, "")

		partialList = mergeUniqueSlices(templateDirPartials, fallbackDirPartials)
	} else {
		partialList = scanPartialsInPath(fallbackDir, "")
	}

	for _, partialName := range partialList {
		logger.LogDebug(fmt.Sprintf("Regist partial %s", partialName))

		bytes, err := ioutil.ReadFile(resolveTemplatePath(partialName)) // just pass the file name
		if err != nil {
			logger.LogError(err)
		}
		partialContent := string(bytes)
		raymond.RegisterPartial(partialName, partialContent)
	}
}

func registerPartialsInPath(path string) {
	logger.LogDebug(fmt.Sprintf("Register partials in path: %s", path))

	files, err := ioutil.ReadDir(path)
	if err != nil {
		logger.LogFatal(err)
	}

	for _, f := range files {
		logger.LogDebug(fmt.Sprintf("%s %t", f.Name(), f.IsDir()))
	}
}

func scanPartialsInPath(path string, prefix string) []string {
	logger.LogDebug(fmt.Sprintf("Scan direcotry: %s", path))

	var filePaths []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		logger.LogFatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			childPaths := scanPartialsInPath(filepath.Join(path, f.Name()), filepath.Join(prefix, f.Name()))
			filePaths = append(filePaths, childPaths...)
		} else {
			filePaths = append(filePaths, filepath.Join(prefix, f.Name()))
		}
	}

	return filePaths
}

func mergeUniqueSlices(first []string, second []string) []string {
	var list []string
	list = append(list, first...)

	sort.Strings(first)
	for _, str := range second {
		if !contains(first, str) {
			list = append(list, str)
		}
	}

	return list
}

func contains(list []string, value string) bool {
	for _, entry := range list {
		if entry == value {
			return true
		}
	}
	return false
}
