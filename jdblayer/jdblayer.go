package jdblayer

import (
	"errors"
	"log"
)

const (
	MONGO = "mongodb"
	MYSQL = "mysql"
	POSQL = "postgres"
)

var errtype = errors.New("Database Type not found... ")

type DBLayer interface {
	AddMember(cm *CrewMember) error
	FindMember(id int) (CrewMember, error)
	AllMembers() (crew, error)
}

type CrewMember struct {
	ID           int    `json:"id" bson:"id"`
	Name         string `json:"name" bson:"Name"`
	SecClearance int    `json:"clearance" bson:"SecurityClearance"`
	Position     string `json:"position" bson:"Title"`
}

type crew []CrewMember

//ConnectDatabase connects to a database type o using the provided connection string
func ConnectDatabase(o string, cstring string) (DBLayer, error) {
	switch o {
	case MONGO:
		return NewMongoStore(cstring)
	case MYSQL:
		return NewMySQLDataStore(cstring)
	case POSQL:
		return NewPgqlDataStore(cstring)
	}
	log.Println("Could not find ", o)
	return nil, errtype
}
