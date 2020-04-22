package getXE_Currency

import (
	"XE_Currency/api/model"
	"XE_Currency/api/utils"
	"fmt"
	_"github.com/lib/pq"
	"log"
	"strconv"
	"strings"
	_"time"
)

var createTableString = `CREATE TABLE IF NOT EXISTS public.exchange_rates(
	       from_currency character varying(3) COLLATE pg_catalog."default",
	       to_currency character varying(3) COLLATE pg_catalog."default",
	       rate numeric,
	       created_at timestamp with time zone,
	   	updated_at timestamp with time zone,
	   	CONSTRAINT unq UNIQUE (from_currency, to_currency)
	   )`

var upsertQuery = `ON CONFLICT (from_currency,to_currency)
				 DO UPDATE
				 SET rate=excluded.rate, updated_at = excluded.updated_at
				 WHERE
				 exchange_rates.from_currency = excluded.from_currency
				 AND
				 exchange_rates.to_currency = excluded.to_currency`

func InitTable() {

	db, err := utils.GetDBConnection()
	if err != nil {
		log.Fatal("Error in GetConnection", err)
		return
	}
	_, err = db.Exec(createTableString)
	if err != nil {
		log.Fatal("Error in creating Table")
		return
	}
	log.Print("Out  table")

	defer db.Close()
	return
}

func BulkInsert(Resp model.XE_Currency_Response) (stmt string, valueArgs []interface{}) {
	valueStrings := make([]string, 0, len(Resp.To))
	valueArgs = make([]interface{}, 0, len(Resp.To)*5)
	for _, to := range Resp.To {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, Resp.From)
		valueArgs = append(valueArgs, to.Quotecurrency)
		valueArgs = append(valueArgs, to.Mid)
		valueArgs = append(valueArgs, Resp.Timestamp)
		valueArgs = append(valueArgs, Resp.Timestamp)
	}
	stmt = fmt.Sprintf("INSERT INTO exchange_rates(from_currency,to_currency,rate,created_at,updated_at) VALUES %s",
		strings.Join(valueStrings, ","))
	stmt = ReplaceSQL(stmt, "(?, ?, ?, ?, ?)", len(valueStrings))
	return
}

func ReplaceSQL(stmt, pattern string, len int) string {
	pattern += ","
	n := 0
	for strings.IndexByte(stmt, '?') != -1 {
		n++
		param := "$" + strconv.Itoa(n)
		stmt = strings.Replace(stmt, "?", param, 1)
	}
	stmt = strings.TrimSuffix(stmt, ",)")
	return stmt
}

func updateDB(Resp model.XE_Currency_Response) (err error) {
	log.Print("In updateDB")
	db, err := utils.GetDBConnection()
	if err != nil {
		log.Fatal("Error in GetConnection", err)
		return
	}
	defer db.Close()

	query, args:= BulkInsert(Resp)

	query += upsertQuery

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("error in prep query: ", err)
		return
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		log.Fatal("error in exec query: ", err)
		return
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("affected = %d\n", rowCnt)
	return
}

