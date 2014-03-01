package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

	"database/sql"
	_ "github.com/lib/pq"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=gopherway sslmode=disable")
	PanicIf(err)
	return db
}

func main() {
	m := martini.Classic()
	m.Map(SetupDB())

	// Serve static files from "static" directory.
	StaticOptions := martini.StaticOptions{Prefix: "static"}
	m.Use(martini.Static("static", StaticOptions))

	// Render html templates from "templates" directory.
	m.Use(render.Renderer())

	m.Get("/", IndexHandler)
	m.Post("/subscribe/", binding.Form(Subscription{}), SubscribeHandler)

	m.Run()
}
