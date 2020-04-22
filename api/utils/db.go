package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
)

func GetDBConnection() (db *sql.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", GetConfig("postgres.host"), GetConfig("postgres.port"), GetConfig("postgres.user"), GetConfig("postgres.password"), GetConfig("postgres.dbname"), GetConfig("postgres.sslmode"))

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
