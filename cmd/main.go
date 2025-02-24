package main

import (
	"log"
	"net/http"
	config "solid-go"
	"solid-go/routes"
	"solid-go/utils"
)

func main() {
	utils.InitENV()
	distFS, err := config.ConfigureFrontend()
	var router *routes.Router

	if utils.ENV == "dev" {
		log.Println("Starting server in development mode...")

		router = routes.InitRouter(distFS)
	} else {
		if err != nil {
			log.Fatalf("failed to create sub filesystem: %v", err)
		}

		router = routes.InitRouter(distFS)
	}

	log.Println("Server started on http://localhost:" + utils.PORT)
	if err := http.ListenAndServe(":"+utils.PORT, router); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
