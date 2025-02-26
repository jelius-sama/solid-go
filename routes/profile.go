package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solid-go/utils"
)

func ProfileRoute(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to HTML
	w.Header().Set("Content-Type", "text/html")
	// Write the HTML content directly to the response
	fmt.Fprint(w, utils.GenerateHTML(utils.HTMLTemplate{
		Body: `
			<div id="root">
				<p id="__server__">Will this get rendered?</p>
			</div>
		`,
		Head: `
			<title id="__server__">Profile - Solid + Go</title>
		`,
	}))
}

func ProfileRouteProps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]any{
		"body": map[string]any{
			"id": "root",
			"children": []map[string]any{
				{"tag": "p", "id": "__server__", "children": "Will this get rendered?"},
			},
		},
		"head": map[string]any{
			"tag":      "title",
			"id":       "__server__",
			"children": "Profile - Solid + Go",
		},
	}
	json.NewEncoder(w).Encode(data)
}
