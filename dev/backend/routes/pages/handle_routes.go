package routes_pages

import (
	"net/http"

	"github.com/Mateus-MS/ProjectZero/dev/backend/app"
	"github.com/Mateus-MS/ProjectZero/dev/backend/middlewares"
)

type RoutesPages struct {
	App *app.Application
}

func RegisterRoutes(app *app.Application) {
	pagesRoutes := &RoutesPages{App: app}

	// Register a clean route
	app.Router.HandleFunc("/test", pagesRoutes.TestPageRoute)

	// Register a route with a middleware
	app.Router.Handle("/test2", middlewares.Chain(
		http.HandlerFunc(pagesRoutes.TestPageRoute),

		middlewares.CorsMiddleware("GET"),
	))
}
