package main

import (
	"bouncedb/http"
	"bouncedb/user"
)

func main() {
	go http.Http()
	user.User()

	select {}
}
