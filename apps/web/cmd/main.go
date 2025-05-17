package main

import "github.com/newspaperhinson/e-commerce-go/apps/web"

func main() {
	app := web.NewApp()
	app.Bootstrap()

	app.Listen(":8080")
}
