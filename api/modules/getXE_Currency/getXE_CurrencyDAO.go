package getXE_Currency

import (
	"XE_Currency/api/model"
	"XE_Currency/api/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var createTableString = `CREATE TABLE IF NOT EXISTS public.exchange_rates(
	       from_currency character varying(3) COLLATE pg_catalog."default",
	       to_currency character varying(3) COLLATE pg_catalog."default",
	       rate double precision,
	       created_at timestamp with time zone,
	   	updated_at timestamp with time zone,
	   	CONSTRAINT unq UNIQUE (from_currency, to_currency)
	   )`


func InitTable() {

	db, err := utils.GetDBConnection()
	if err != nil {
		log.Fatal("Error in GetConnection", err)
		return
	}
	_,err = db.Exec(createTableString)
	if err != nil {
		log.Fatal("Error in creating Table")
		return
	}

	defer db.Close()
	return
}

func insertDB(Resp model.XE_Currency_Response) (err error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		log.Fatal("Error in GetConnection", err)
		return
	}
	defer db.Close()

	query, args := BulkInsert(Resp)

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("error in prep query: ", err)
		return
	}

	//Resp.From,Resp.To[0].Quotecurrency,Resp.To[0].Mid,Resp.Timestamp,time.Now().String()
	res, err := stmt.Exec(args...)
	if err != nil {
		log.Fatal("error in exec query: ", err)
		return
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("affected = %d\n", rowCnt)
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
		valueArgs = append(valueArgs, time.Now().String())
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
	query, args1, args2, args3 := BulkUpdate(Resp)
	log.Print("args3", args1, args2, args3)
	// args4 := []float64{}
	// for _, v := range args3 {
	//  args4 = append(args4, v.(float64))
	// }
	// query := "UPDATE exchange_rates SET rate = $3, created_at = $4, updated_at = $5 Where from_currency = $1 AND to_currency =$2"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("error in prep query: ", err)
		return
	}
	//Resp.From,Resp.To[0].Quotecurrency,Resp.To[0].Mid,Resp.Timestamp,time.Now().String()
	res, err := stmt.Exec(args1, args3)
	if err != nil {
		log.Fatal("error in exec query: ", err)
		return
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("affected = %d\n", rowCnt)
	return
}

func BulkUpdate(Resp model.XE_Currency_Response) (stmt string, valueArgs1, valueArgs2, valueArgs3 []interface{}) {
	valueArgs1 = make([]interface{}, 0, 1)
	valueArgs2 = make([]interface{}, 0, len(Resp.To))
	valueArgs3 = make([]interface{}, 0, len(Resp.To))
	valueArgs1 = append(valueArgs1, Resp.From)
	for _, to := range Resp.To {
		valueArgs2 = append(valueArgs2, to.Quotecurrency)
		valueArgs3 = append(valueArgs3, to.Mid)
	}

	stmt = "update exchange_rates set rate = data_table.rate from (select $1 as from_currency, $2 as rate) as data_table where exchange_rates.from_currency = data_table.from_currency AND exchange_rates.to_currency = data_table.to_currency"
	return
}
