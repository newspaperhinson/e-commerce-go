package web

import (
	"fmt"
	"net/http"
	"strconv"
)

type App struct{}

func NewApp() *App {
	app := App{}
	app.Bootstrap()

	return &app
}

func (a *App) Bootstrap() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}

func (a *App) Listen(port int) {
	fmt.Printf("Listening on port %d...\n", port)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
