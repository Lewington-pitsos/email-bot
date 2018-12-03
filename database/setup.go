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
	_, err := s.db.Exec(
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

	generalhelpers.Check(err)

	_, err2 := s.db.Exec(
		`CREATE TABLE yandex_profiles (
			id serial,
			first_name varchar(30) NOT NULL,
			last_name varchar(30) NOT NULL,
			password varchar(40) NOT NULL,
			email varchar(40) NOT NULL,
			question varchar(60) NOT NULL,
			answer varchar(30) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW() NOT NULL,
			PRIMARY KEY(id) 
		);`,
	)

	generalhelpers.Check(err2)
}

func (s *Setup) Drop() {
	s.db.Exec("DROP TABLE profiles")
	s.db.Exec("DROP TABLE yandex_profiles")
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
