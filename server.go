package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"net/http"
)

func SubscribeHandler(rnd render.Render, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	print(name)

	if email == "vaxxxa@gmail.com" {
		rnd.JSON(200, map[string]string{"status": "ok"})
	} else {
		rnd.JSON(200, map[string]string{"status": "fail"})
	}
}

func IndexHandler(r render.Render) {
	r.HTML(200, "index", nil)
}

func main() {
	m := martini.Classic()

	// Serve static files from "static" directory.
	StaticOptions := martini.StaticOptions{Prefix: "static"}
	m.Use(martini.Static("static", StaticOptions))

	// Render html templates from "templates" directory.
	m.Use(render.Renderer())

	m.Get("/", IndexHandler)
	m.Post("/subscribe/", SubscribeHandler)

	m.Run()
}
