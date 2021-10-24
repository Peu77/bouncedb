package main

import (
	"bouncedb/config"
	"bouncedb/database"
	"bouncedb/file"
	"bouncedb/http"
	"bouncedb/user"
	"bouncedb/user/auth"
)

func main() {
	file.InitFiles()
	database.InitDatabases()
	config.InitConfig()
	go http.Http()
	user.User()

	auth.CreateToken([]string{"database.conquest.*", "database.conquest.read", "database.conquest.write"})

	select {}
}
