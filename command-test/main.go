package main

import (
	"command-ui/components"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {

	home := &components.App{}
	app.Route("/", home)

	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "test",
		Styles: []string{
			"/web/style.css",
		},
	})
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}

}
