package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "order_assignment"
)

func ConnectDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Error opening database: %v", err)
	}

	fmt.Println("Database connected!")
	return db
}
