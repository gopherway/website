package main

import (
	"github.com/martini-contrib/binding"

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
