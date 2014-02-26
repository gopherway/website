package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	// Serve static files from "static" directory.

	// Render html templates from "templates" directory.
	StaticOptions := martini.StaticOptions{Prefix: "static"}
	m.Use(martini.Static("static", StaticOptions))
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", "")
	})

	m.Use(martini.Static("static"))

	m.Run()
}
