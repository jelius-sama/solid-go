package routes

import (
	"encoding/json"
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
			<div id="root">
			<ol id="__server__">
				<li>Apples</li>
				<li>Bananas</li>
				<li>Strawberries</li>
				<li>Pinapples</li>
			</ol>
			</div>
		`,
		Head: `
			<title id="__server__">Settings - Solid + Go</title>
		`,
	}))
}

func SettingsRouteProps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]any{
		"body": map[string]any{
			"id": "root",
			"children": []map[string]any{
				{"tag": "ol", "id": "__server__", "children": []string{"Apples", "Bananas", "Strawberries", "Pineapples"}},
			},
		},
		"head": map[string]any{
			"tag":      "title",
			"id":       "__server__",
			"children": "Settings - Solid + Go",
		},
	}
	json.NewEncoder(w).Encode(data)
}
