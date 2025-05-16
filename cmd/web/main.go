package main

import "github.com/newspaperhinson/e-commerce-go/internal/web"

func main() {
	app := web.NewApp()

	web.Serve(app, 8080)
}
