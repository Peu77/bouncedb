package main

import (
	"bouncedb/format"
	"bouncedb/http"
	"bouncedb/user"
)

func main() {
	format.Test()
	go http.Http()
	user.User()

	select {}
}
