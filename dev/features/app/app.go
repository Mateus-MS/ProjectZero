package projectzero_app

import (
	"database/sql"
	"net/http"
	"sync"
)

var app_instance *Application
var app_once sync.Once

type Application struct {
	DB     *sql.DB
	Router *http.ServeMux
}

func GetInstance() *Application {
	app_once.Do(func() {
		app_instance = newApplication()
	})
	return app_instance
}

func newApplication() *Application {
	// Create the router
	router := http.NewServeMux()

	// Serve static files from the "frontend" directory
	router.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("dev/frontend"))))

	// Return the application instance
	return &Application{
		// DB:     GetInstance(),
		Router: router,
	}
}

func (app *Application) RegisterRoutes(route string, handler http.HandlerFunc) {
	app.Router.HandleFunc(route, handler)
}
