package routes_pages

import (
	"net/http"
	"strings"

	test_page_mob "github.com/Mateus-MS/ProjectZero/dev/frontend/mobile/pages/test_page"
)

func (app *RoutesPages) TestPageRoute(w http.ResponseWriter, r *http.Request) {
	// Send the page only if the user agent is mobile
	if strings.Contains(r.UserAgent(), "Mobile") {
		test_page_mob.TestPage("test").Render(r.Context(), w)
	}

	// If the user agent is not mobile, return a 404 error
	http.Error(w, "Page not found", http.StatusNotFound)
}
