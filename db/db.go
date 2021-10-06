package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fatih/color"

	_ "github.com/lib/pq"
)

const (
    HOST = "database"
    PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
    Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
    
    color.Blue("YO")
    log.Println(username)
    log.Println(password)
    log.Println(database)

    db := Database{}
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        HOST, PORT, username, password, database)
    conn, err := sql.Open("postgres", dsn)
    if err != nil {
        return db, err
    }
    db.Conn = conn
    err = db.Conn.Ping()
    if err != nil {
        return db, err
    }
    log.Println("Database connection established")
    return db, nil
}