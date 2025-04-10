package routes_pages

import "net/http"

func (app *RoutesPages) LearnPageRoute(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Test page</title>
		</head>
		<body>
			<h1>Welcome to test page</h1>
			<p>I'm just testing</p>
		</body>
		</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
