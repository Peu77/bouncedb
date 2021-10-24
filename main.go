package main

import (
	"bouncedb/config"
	"bouncedb/database"
	"bouncedb/file"
	"bouncedb/http"
	"bouncedb/user"
	"fmt"
)

func main() {
	file.InitFiles()
	config.InitConfig()
	database.InitDatabases()
	go http.Http()
	user.User()
	fmt.Println()

	select {}
}
