package db

import (
	"database/sql"
	"fmt"
	"xe-currency/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
)

const (
	dbdriver = "postgres"
)

func dbInit() (db *sql.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.GetConfig("postgres.host"),
		config.GetConfig("postgres.port"),
		config.GetConfig("postgres.user"),
		config.GetConfig("postgres.password"),
		config.GetConfig("postgres.dbname"),
		config.GetConfig("postgres.sslmode"))

	db, err = sql.Open(dbdriver, dsn)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot initialize database")
		return
	}
	logger.WithField("dsn", dsn).Info("DB connected Successfully")
	return
}

func ExecQuery(query string, args []interface{}) (result sql.Result, err error) {
	db, err := dbInit()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot initialize database")
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		logger.WithField("error in prep query", err.Error()).Error("Query Failed")
		return
	}

	result, err = stmt.Exec(args...)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	return
}
