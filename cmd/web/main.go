package main

import "github.com/newspaperhinson/e-commerce-go/internal/web"

func main() {
	app := web.NewApp()

	app.Listen(8080)
}
