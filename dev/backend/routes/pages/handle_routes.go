package routes_pages

import "github.com/Mateus-MS/ProjectZero.git/dev/backend/app"

type RoutesPages struct {
	App *app.Application
}

func RegisterRoutes(app *app.Application) {
	pagesRoutes := &RoutesPages{App: app}

	app.Router.HandleFunc("/test", pagesRoutes.TestPageRoute)
}
