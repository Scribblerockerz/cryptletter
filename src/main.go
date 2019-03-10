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

	// TODO: implement dynamic partial handling
	// - scan template dir for all sub dirs and register them as partials
	// - scan fallback dir and register all missing / not declared templates

	// TODO: intergrate redis server
	// - connect to server
	// - set keys with default ttl of 30 days
	// - get keys and match them against users ip
	// 	- if no ip is set
	//		- update ttl to requested value

	router := NewRouter()
	port := fmt.Sprintf("%d", Config.Server.Port)

	fmt.Printf("Serving taskronaut api on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func prepareTemplates() {
	raymond.RegisterPartial("foo", "<strong>FOO</strong>")
}
