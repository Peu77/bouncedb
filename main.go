package main

import (
	"bouncedb/format"
	"bouncedb/http"
	"bouncedb/user"
)

func main() {
	format.Test()
	http.Http()
	user.User()
}
