package web

import "net/http"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Bootstrap() {
}

func (a *App) Listen(address string) {
	http.ListenAndServe(address, nil)
}
