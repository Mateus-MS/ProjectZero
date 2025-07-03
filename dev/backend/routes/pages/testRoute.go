package routes_pages

import (
	"net/http"
	projectzero_app "placeholder/dev/features/app"
	projectzero_middlewares "placeholder/dev/features/middlewares"
	test_page_mob "placeholder/dev/frontend/desktop/pages/test_page"
	"strings"
)

func init() {
	projectzero_app.GetInstance().RegisterRoutes(
		"/test/route",
		projectzero_middlewares.Chain(
			http.HandlerFunc(TestPageRoute),
			projectzero_middlewares.CorsMiddleware("GET"),
		).ServeHTTP,
	)
}

func TestPageRoute(w http.ResponseWriter, r *http.Request) {
	// Send the page only if the user agent is mobile
	if strings.Contains(r.UserAgent(), "Mobile") {
		test_page_mob.TestPage("test").Render(r.Context(), w)
	}

	// If the user agent is not mobile, return a 404 error
	http.Error(w, "Page not found", http.StatusNotFound)
}
