package main

import (
	"bouncedb/database"
	"bouncedb/file"
	"bouncedb/http"
	"bouncedb/user"
)

func main() {
	file.InitFiles()
	database.InitDatabases()
	go http.Http()
	user.User()

	select {}
}
