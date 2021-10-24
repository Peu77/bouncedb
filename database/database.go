package database

import "github.com/google/uuid"

type Database struct {
	Name string
	Id   uuid.UUID
	sets []Set
}
