package database

import "database/sql"

const connStr string = "user=email_bot dbname=email_bot password=beepBoop sslmode=disable"
const driverName string = "postgres"

type connection struct {
	db *sql.DB
}

func (c *connection) Close() {
	c.db.Close()
}
