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