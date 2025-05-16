package web

import (
	"fmt"
	"net/http"
	"strconv"
)

type App struct{}

func NewApp() *App {
	return Bootstrap(&App{})
}

func Bootstrap(app *App) *App {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	return app
}

func Serve(app *App, port int) {
	fmt.Printf("Listening on port %d...\n", port)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
