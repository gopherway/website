package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"fmt"
	"net/http"
)

func SubscribeHandler(r *http.Request) string {
	fmt.Println(r.FormValue("name"))
	fmt.Println(r.FormValue("email"))
	return "1"
}

func IndexHandler(r render.Render) {
	r.HTML(200, "index", nil)
}

func main() {
	m := martini.Classic()

	// Serve static files from "static" directory.

	// Render html templates from "templates" directory.
	StaticOptions := martini.StaticOptions{Prefix: "static"}
	m.Use(martini.Static("static", StaticOptions))
	m.Use(render.Renderer())

	m.Get("/", IndexHandler)
	m.Post("/subscribe/", SubscribeHandler)

	m.Run()
}
