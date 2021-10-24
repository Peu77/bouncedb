package database

import (
	"bouncedb/utils"
	"github.com/google/uuid"
)

type Database struct {
	Name string
	Id   uuid.UUID
	sets []Set
}

func NewDatabase(name string) Database {
	database := Database{}
	database.Name = name
	database.Id = utils.NewId()
	database.sets = []Set{}
	return database
}
