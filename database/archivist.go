package database

import (
	"database/sql"
	"email-bot/helpers/generalhelpers"

	_ "github.com/lib/pq"
)

type Archivist struct {
	connection
}

func NewArchivist() *Archivist {
	database, err := sql.Open(driverName, connStr)
	generalhelpers.Check(err)
	return &Archivist{
		connection{
			db: database,
		},
	}
}
