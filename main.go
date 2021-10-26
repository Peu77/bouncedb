package main

import (
	"bouncedb/config"
	"bouncedb/database"
	"bouncedb/file"
	"bouncedb/http"
	"bouncedb/user"
	"bouncedb/user/auth"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	file.InitFiles()
	database.InitDatabases()
	config.InitConfig()
	go http.Http()
	user.User()

	auth.CreateToken([]string{"database.conquest.*", "database.conquest.read", "database.conquest.write"})

	created := database.CreateDatabase(database.NewDatabase("test"))

	fmt.Println(created)

	deleted := database.DeleteDatabase(uuid.MustParse("7030e915-3684-11ec-9c08-309c23168291"))
	fmt.Println(deleted)
	select {}
}
