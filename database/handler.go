package database

import (
	"bouncedb/file"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

var Databases []Database

const configFile = "files/databases.conf"

func InitDatabases() {
	if file.ExistFile(configFile) {
		jsonString := file.ReadFile(configFile)
		err := json.Unmarshal([]byte(jsonString), &Databases)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(Databases)
	} else {
		newUUID, err := uuid.NewUUID()
		if err != nil {
			return
		}
		Databases = append(Databases, Database{"stats", newUUID, []Set{}})
		marshal, err := json.MarshalIndent(Databases, "", "  ")
		if err != nil {
			return
		}
		file.WriteInFile(configFile, string(marshal))
	}
}
