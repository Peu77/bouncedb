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
	} else {
		SaveConfig()
	}
}

// SaveConfig overwrite the content of configFile with Databases array
func SaveConfig() {
	marshal, err := json.MarshalIndent(Databases, "", "  ")
	if err != nil {
		return
	}
	file.WriteInFile(configFile, string(marshal))
}

func CreateDatabase(database Database) bool {
	if !existName(database.Name) {
		Databases = append(Databases, database)
		file.Mkdir(database.GetPath())
		SaveConfig()
		return true
	}
	return false
}

func DeleteDatabase(id uuid.UUID) bool {
	index := findElement(id)
	if index != -1 {
		file.RmDir(Databases[index].GetPath())
		Databases = RemoveIndex(Databases, index)
		SaveConfig()
		return true
	}
	return false
}

func RemoveIndex(s []Database, index int) []Database {
	return append(s[:index], s[index+1:]...)
}

func findElement(id uuid.UUID) int {
	for i, database := range Databases {
		if database.Id == id {
			return i
		}
	}
	return -1
}

func GetDatabaseName(name string) (*Database, int) {
	for i, database := range Databases {
		if database.Name == name {
			return &Databases[i], 0
		}
	}
	return &Database{}, 1
}

func existName(name string) bool {
	_, error := GetDatabaseName(name)
	return error == 0
}
