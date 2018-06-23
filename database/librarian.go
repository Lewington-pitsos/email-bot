package database

import (
	"database/sql"
	"email-bot/helpers/generalhelpers"

	_ "github.com/lib/pq"
)

type Librarian struct {
	connection
}

func (l *Librarian) profileQuery() *sql.Stmt {
	stmt, err := l.db.Prepare("SELECT first_name, last_name, pass, email FROM profiles LIMIT $1;")

	generalhelpers.Check(err)
	return stmt
}

func (l *Librarian) allProfiles(stmt *sql.Stmt, number int) []map[string]string {
	rows, err := stmt.Query(number)
	generalhelpers.Check(err)

	profiles := make([]map[string]string, 0, number)

	var firstName, lastName, pass, email string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&firstName, &lastName, &pass, &email)
		generalhelpers.Check(err)

		profile := map[string]string{
			"firstname": firstName,
			"lastname":  lastName,
			"pass":      pass,
			"email":     email,
		}

		profiles = append(profiles, profile)
	}

	err = rows.Err()
	generalhelpers.Check(err)

	return profiles
}

func (l *Librarian) FindProfiles(number int) []map[string]string {
	stmt := l.profileQuery()
	profiles := l.allProfiles(stmt, number)
	return profiles
}

func NewLibrarian() *Librarian {
	database, err := sql.Open(driverName, connStr)
	generalhelpers.Check(err)
	return &Librarian{
		connection{
			db: database,
		},
	}
}
