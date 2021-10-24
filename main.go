package main

import (
	"bouncedb/file"
	"bouncedb/http"
	"bouncedb/user"
)

func main() {
	file.InitFiles()
	go http.Http()
	user.User()

	select {}
}
