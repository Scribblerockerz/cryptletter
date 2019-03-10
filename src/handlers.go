package main

import (
	"fmt"
	"net/http"

	"github.com/aymerick/raymond"
)

// Index actions
func Index(w http.ResponseWriter, r *http.Request) {

	result := RenderLayout("src/templates/layout.hbs", map[string]string{
		"title": "My New Post!",
		"body": RenderLayout("src/templates/index.hbs", map[string]string{
			"headline": "Can You Make More Money With A Mobile App Or A PWA?",
			"content":  " Take a gander at the revenues of the top mobile apps and it’s easy to get lost in dreams of what could be if only you built a mobile app today. Then again, have you ever considered how much it actually costs to build and maintain a mobile app? When you look at the big picture, you’ll soon realize that mobile apps aren’t a smart investment for most. That’s why you need to give serious consideration to building a PWA this year.",
		}),
	})

	fmt.Fprintf(w, result)
}

// RenderLayout renders a layout path with a context
func RenderLayout(path string, ctx map[string]string) string {
	tpl, err := raymond.ParseFile(path)
	if err != nil {
		panic(err)
	}

	result := tpl.MustExec(ctx)
	return result
}
