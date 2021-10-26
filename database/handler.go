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
		saveConfig()
	}
}

// overwrite the content of configFile with Databases array
func saveConfig() {
	marshal, err := json.MarshalIndent(Databases, "", "  ")
	if err != nil {
		return
	}
	file.WriteInFile(configFile, string(marshal))
}

func CreateDatabase(database Database) bool {
	if !existName(database.Name) {
		Databases = append(Databases, database)
		file.Mkdir(database.getPath())
		saveConfig()
		return true
	}
	return false
}

func DeleteDatabase(id uuid.UUID) bool {
	index := findElement(id)
	if index != -1 {
		file.RmDir(getElement(index).getPath())
		Databases = RemoveIndex(Databases, index)
		saveConfig()
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

func getElement(index int) Database {
	for i, database := range Databases {
		if i == index {
			return database
		}
	}
	return Database{}
}

func existName(name string) bool {
	for _, database := range Databases {
		if database.Name == name {
			return true
		}
	}
	return false
}
