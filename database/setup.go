package database

import (
	"database/sql"
	"email-bot/helpers/generalhelpers"

	_ "github.com/lib/pq"
)

type Setup struct {
	connection
}

func (s *Setup) Setup() {
	s.db.Exec(
		`CREATE TABLE profiles (
			id serial,
			first_name varchar(30) NOT NULL,
			last_name varchar(30) NOT NULL,
			pass varchar(40) NOT NULL,
			email varchar(40) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW() NOT NULL,
			birthdate TIMESTAMP NOT NULL,
			PRIMARY KEY(id) 
		);`,
	)
}

func (s *Setup) Drop() {
	s.db.Exec("DROP TABLE profiles")
}

func NewSetup() *Setup {
	database, err := sql.Open(driverName, connStr)
	generalhelpers.Check(err)
	return &Setup{
		connection{
			db: database,
		},
	}
}
