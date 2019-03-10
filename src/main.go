package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aymerick/raymond"
)

func main() {
	AssembleConfiguration()
	prepareTemplates()

	router := NewRouter()
	port := fmt.Sprintf("%d", Config.Server.Port)

	fmt.Printf("Serving taskronaut api on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func prepareTemplates() {
	raymond.RegisterPartial("foo", "<strong>FOO</strong>")
}
