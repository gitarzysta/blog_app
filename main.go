package main

import (
	"log"

	"github.com/gitarzysta/blog_app/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
