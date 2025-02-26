package routes

import (
	"io/fs"
	"net/http"
	"strings"
)

type Router struct {
	routes     map[string]http.HandlerFunc
	fileServer http.Handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Check if there's an exact route match
	if handler, exists := r.routes[req.URL.Path]; exists {
		handler(w, req)
		return
	}

	// Handle static files (for `/assets/*`)
	if strings.HasPrefix(req.URL.Path, "/assets/") {
		r.fileServer.ServeHTTP(w, req)
		return
	}

	// If no match, return 404
	NotFoundHandler(w, req)
}

// Register route
func (r *Router) Handle(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

func InitRouter(distFS fs.FS) *Router {
	router := &Router{
		routes:     make(map[string]http.HandlerFunc),
		fileServer: http.FileServer(http.FS(distFS)),
	}

	router.Handle("/", HomeRoute)
	router.Handle("/profile", ProfileRoute)
	router.Handle("/profile/props", ProfileRouteProps)
	router.Handle("/settings", SettingsRoute)
	router.Handle("/settings/props", SettingsRouteProps)

	return router
}
