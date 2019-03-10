package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aymerick/raymond"
)

func main() {

	config := AssembleConfiguration()

	router := NewRouter()
	port := fmt.Sprintf("%d", config.Server.Port)

	prepareTemplates()

	fmt.Printf("Serving taskronaut api on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func prepareTemplates() {
	raymond.RegisterPartial("foo", "<strong>FOO</strong>")
}
