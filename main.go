package main

import "github.com/daddydemir/got/handlers"

func main() {

	app := handlers.Urls()
	_ = app.Listen(":1345")
}
