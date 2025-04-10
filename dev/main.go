package main

import (
	"net/http"

	"github.com/Mateus-MS/ProjectZero.git/dev/backend/app"
	routes_pages "github.com/Mateus-MS/ProjectZero.git/dev/backend/routes/pages"
)

func main() {
	app := app.NewApplication()

	routes_pages.RegisterRoutes(app)

	startServer(app.Router, "dev")
}

func startServer(router *http.ServeMux, env string) {
	if env == "dev" {
		println("Starting SERVER in DEV mode")
		go func() {
			if err := http.ListenAndServe("localhost:3000", router); err != nil {
				println("HTTP server error:", err)
			}
		}()
		select {}
	}
}
