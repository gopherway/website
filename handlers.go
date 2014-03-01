package main

import (
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

	"database/sql"
	"strings"
)

func SubscribeHandler(rnd render.Render, formErr binding.Errors, sub Subscription, db *sql.DB) {
	if formErr.Count() == 0 {
		rows, err := db.Query(`INSERT INTO subscriptions
			(name, email) VALUES ($1, $2)`, sub.Name, sub.Email)

		PanicIf(err)
		defer rows.Close()

		rnd.JSON(200, map[string]string{"status": "ok"})
	} else {
		var errors []string
		for _, value := range formErr.Fields {
			errors = append(errors, value)
		}

		strErrors := strings.Join(errors, ", ")
		rnd.JSON(200, map[string]string{"status": "fail", "errors": strErrors})
	}
}

func IndexHandler(r render.Render) {
	r.HTML(200, "index", nil)
}
