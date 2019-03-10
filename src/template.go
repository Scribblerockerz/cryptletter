package main

import (
	"os"

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
	templateDir := Config.App.TemplatesDir
	fallbackDir := DefaultConfiguration.App.TemplatesDir

	if _, err := os.Stat(templateDir + "/" + path); err == nil {
		return templateDir + "/" + path

	} else if os.IsNotExist(err) {
		return fallbackDir + "/" + path
	} else {
		panic(err)
	}
}
