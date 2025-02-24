package routes

import (
	"fmt"
	"net/http"
	"solid-go/utils"
)

func SettingsRoute(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to HTML
	w.Header().Set("Content-Type", "text/html")
	// Write the HTML content directly to the response
	fmt.Fprint(w, utils.GenerateHTML(utils.HTMLTemplate{
		Body: `
			<div id="root"></div>
		`,
		Head: `
			<title>Settings | Solid + Go</title>
		`,
	}))
}
