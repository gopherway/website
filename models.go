package main

import (
	"github.com/martini-contrib/binding"

	"database/sql"
	"net/http"
	"regexp"
)

type Subscription struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func (sub Subscription) Validate(errors *binding.Errors, req *http.Request) {
	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(sub.Email))
	if matched == false {
		errors.Fields["email"] = "Please enter a valid email address."
	}
}

func (sub Subscription) Exist(db *sql.DB) bool {
	var result int

	row := db.QueryRow("SELECT 1 from subscriptions WHERE email = $1", sub.Email)
	err := row.Scan(&result)
	return err == nil
}

func (sub Subscription) Save(db *sql.DB) {
	rows, err := db.Query(`INSERT INTO subscriptions
			(name, email) VALUES ($1, $2)`, sub.Name, sub.Email)
	PanicIf(err)
	defer rows.Close()
}
