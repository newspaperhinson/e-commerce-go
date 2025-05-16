package web

import (
	"fmt"
	"net/http"
	"strconv"
)

type App struct{}

func Bootstrap() {
	app := &App{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	serve(app, 8080)
}

func serve(app *App, port int) {
	fmt.Printf("Listening on port %d...\n", port)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
