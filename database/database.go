package database

import (
	"bouncedb/utils"
	"github.com/google/uuid"
)

type Database struct {
	Name string
	Id   uuid.UUID
	Sets []Set
}

func NewDatabase(name string) Database {
	database := Database{}
	database.Name = name
	database.Id = utils.NewId()
	database.Sets = []Set{}
	return database
}

func (database Database) GetPath() string {
	return "files/databases/" + database.Name
}

func (database *Database) CreateSet(name string) bool {
	if !database.existSetName(name) {
		set := Set{name, utils.NewId()}
		database.Sets = append(database.Sets, set)
		SaveConfig()
		return true
	}
	return false
}

func (database *Database) DeleteSet(id uuid.UUID) bool {
	index := database.findSetElement(id)
	if index != -1 {
		database.Sets = removeIndex(database.Sets, index)
		SaveConfig()
		return true
	}
	return false
}

func (database Database) existSetName(name string) bool {
	for _, set := range database.Sets {
		if set.Name == name {
			return true
		}
	}
	return false
}

func (database Database) findSetElement(id uuid.UUID) int {
	for i, set := range database.Sets {
		if set.Id == id {
			return i
		}
	}
	return -1
}

func removeIndex(s []Set, index int) []Set {
	return append(s[:index], s[index+1:]...)
}
