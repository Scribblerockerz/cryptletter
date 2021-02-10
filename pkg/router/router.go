package router

import (
	"github.com/Scribblerockerz/cryptletter/pkg/handler"
	"github.com/Scribblerockerz/cryptletter/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes slice
type Routes []Route

const defaultStaticDirPathPrefix = "/s/"
const staticDirPathPrefix = "/static/"
const defaultAssetDir = "web/public"
const assetDir = "web/public"

// NewRouter factory
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var h http.Handler

		h = route.HandlerFunc
		h = logger.HTTPLogger(h, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(h)
	}

	router.NotFoundHandler = logger.HTTPLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.NotFound(w, r)
	}), "404")

	// Handle static assets
	router.
		PathPrefix(defaultStaticDirPathPrefix).
		Handler(http.StripPrefix(defaultStaticDirPathPrefix, http.FileServer(http.Dir(defaultAssetDir))))

	// Override default assets by placing them into the second dir
	router.
		PathPrefix(staticDirPathPrefix).
		Handler(http.StripPrefix(staticDirPathPrefix, http.FileServer(http.Dir(assetDir))))

	return router
}

var routes = Routes{
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: handler.IndexAction,
	},
	//Route{
	//	Name:        "Styleguide",
	//	Method:      "GET",
	//	Pattern:     "/styleguide",
	//	HandlerFunc: handler.StyleguideAction,
	//},
	Route{
		Name:        "NewMessage",
		Method:      "POST",
		Pattern:     "/",
		HandlerFunc: handler.NewMessageAction,
	},
	Route{
		Name:        "ShowMessage",
		Method:      "GET",
		Pattern:     "/{token}/",
		HandlerFunc: handler.ShowAction,
	},
	Route{
		Name:        "DeleteMessage",
		Method:      "DELETE",
		Pattern:     "/{token}/",
		HandlerFunc: handler.DeleteMessageAction,
	},
}
