package main

import (
	"github.com/codegangsta/martini"
	flags "github.com/jessevdk/go-flags"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"os"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	Author  = "Vasyl Nakvasiuk"
	Version = "0.1"
)

var opts struct {
	StaticDir    string `short:"s" long:"static" description:"path to static directory" default:"static"`
	TemplatesDir string `short:"t" long:"templates" description:"path to templates directory" default:"templates"`
}

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
	_, err := flags.ParseArgs(&opts, os.Args)
	PanicIf(err)

	m := martini.Classic()
	m.Map(SetupDB())

	// Serve static files from "static" directory.
	StaticOptions := martini.StaticOptions{Prefix: "static"}
	m.Use(martini.Static(opts.StaticDir, StaticOptions))

	// Render html templates from "templates" directory.
	m.Use(render.Renderer(render.Options{Directory: opts.TemplatesDir}))

	m.Get("/", IndexHandler)
	m.Post("/subscribe/", binding.Form(Subscription{}), SubscribeHandler)

	m.Run()
}
