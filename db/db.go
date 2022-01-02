package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// HOST = "database"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

// Connect to the Database
func Initialize(username, password, database string, host string) (Database, error) {
	db := Database{}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, PORT, username, password, database,
	)

	conn, err := sql.Open("postgres", dsn)
	
	if err != nil {
		return db, err
	}
	
	db.Conn = conn
	err = db.Conn.Ping()

	if err != nil {
		return db, err
	}

	dbConnectionSuccess()
	return db, nil
}

func dbConnectionSuccess() {
	log.Println("------------------------------------------------------")
	log.Println("------------DATABASE CONNECTION ESTABLISHED-----------")
	log.Println("------------------------------------------------------")
}