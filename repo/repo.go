package repo

import (
	"database/sql"
	"log"

	"github.com/cseltol/test_quadro/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	conn, err := sql.Open("mysql", config.DATABASE_URL)
	if err != nil {
		log.Fatalf("failed to connect to DB, cause: %v\n", err)
		return nil
	}
	return conn
}

func InitDB() {
	c := GetConnection()
	defer c.Close()

	c.Exec(
		"CREATE DATABASE bookshelf;",
	)
	c.Exec(
		"\\c bookshelf",
	)
	c.Exec(
		`CREATE TABLE IF NOT EXISTS books (
			id serial PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			author VARCHAR(50) NOT NULL
		);`,
	)
	c.Exec(
		`CREATE TABLE IF NOT EXISTS authors (
			id serial PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		);`,
	)
}