package main

import (
	"flag"
	"net/http"
	_ "placeholder/dev/backend/routes/pages"
	app "placeholder/dev/features/app"
)

func main() {
	enviroment := flag.String("env", "dev", "The enviroment to run")

	flag.Parse()

	app := app.GetInstance()

	startServer(app.Router, *enviroment)
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
