package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// HOST = "host.docker.internal"
	// HOST = "localhost"
	HOST = "database"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {

	log.Println("Host --> ", HOST)
	log.Println("Port --> ", PORT)
	log.Println("Username --> ", username)
	log.Println("Password --> ", password)
	log.Println("Database --> ", database)

	db := Database{}

	log.Println("DB? --> ", db)

	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	"localhost", 5432, "postgres", "postgres", "postgres")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)

	log.Println("DSN --> ", dsn)
	log.Println("--------")

	conn, err := sql.Open("postgres", dsn)
	log.Println(conn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("-----------DATABASE CONNECTION ESTABLISHED-----------")
	return db, nil
}
