package main

import (
	"github.com/martini-contrib/render"

	"database/sql"
	"strings"
)

func SubscribeHandler(rnd render.Render, sub Subscription, db *sql.DB) {
	var errors []string

	if sub.Exist(db) {
		errors = append(errors, "This email is already subscribed.")
	}

	if len(errors) == 0 {
		sub.Save(db)

		rnd.JSON(200, map[string]string{"status": "ok"})
	} else {
		strErrors := strings.Join(errors, ", ")
		rnd.JSON(200, map[string]string{"status": "fail", "errors": strErrors})
	}
}

func IndexHandler(r render.Render) {
	r.HTML(200, "index", nil)
}
