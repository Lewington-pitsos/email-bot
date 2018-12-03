package database

import (
	"database/sql"
	"email-bot/helpers/generalhelpers"

	_ "github.com/lib/pq"
)

type Archivist struct {
	connection
}

// func (a *Archivist) profileRecordPS() *sql.Stmt {
// 	stmt, err := a.db.Prepare(`INSERT INTO profiles (
// 		first_name,
// 		last_name,
// 		pass,
// 		email,
// 		birthdate
// 	) VALUES (
// 		$1,
// 		$2,
// 		$3,
// 		$4,
// 		$5
// 	)`)

// 	generalhelpers.Check(err)
// 	return stmt
// }

// func (a *Archivist) RecordProfile(profile map[string]string) {
// 	stmt := a.profileRecordPS()
// 	_, err := stmt.Exec(
// 		profile["firstname"],
// 		profile["lastname"],
// 		profile["pass"],
// 		profile["email"],
// 		profile["birthdate"],
// 	)
// 	generalhelpers.Check(err)
// }

func (a *Archivist) yandexProfileRecordPS() *sql.Stmt {
	stmt, err := a.db.Prepare(`INSERT INTO yandex_profiles (
		first_name, 
		last_name, 
		password, 
		email, 
		question,
		answer
	) VALUES (
		$1,
		$2,
		$3, 
		$4, 
		$5,
		$6
	)`)

	generalhelpers.Check(err)
	return stmt
}

func (a *Archivist) RecordYandexProfile(profile map[string]string) {
	stmt := a.yandexProfileRecordPS()
	_, err := stmt.Exec(
		profile["firstname"],
		profile["lastname"],
		profile["password"],
		profile["email"],
		profile["question"],
		profile["answer"],
	)
	generalhelpers.Check(err)
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
