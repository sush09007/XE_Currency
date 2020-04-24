package service

import (
	"fmt"
	"strconv"
	"xe-currency/config"
	"xe-currency/db"
	"xe-currency/model"

	"strings"
	"sync"
	"time"

	logger "github.com/sirupsen/logrus"
)

func InitJob() int {

	start := time.Now()
	var currencies = config.GetStringSlice("currency")

	logger.WithField("currencies", currencies).Info("Currencies Initialized")

	var w sync.WaitGroup
	var to = strings.Join(currencies, ",")

	for _, c := range currencies {
		w.Add(1)
		go job(c, to, &w)
	}

	w.Wait()

	elapsed := time.Now().Sub(start)
	logger.WithField("Total time taken:", elapsed).Info("Exec Time")

	return 1
}

func job(from, to string, wg *sync.WaitGroup) {

	defer wg.Done()

	resp, err := httpReqToXE(to, from)
	if err != nil {
		return
	}
	_ = resp

	query, args := updateQuery(resp)

	result, err := db.ExecQuery(query, args)
	if err != nil {
		logger.WithField("error in db update", err.Error()).Error("Update DB Failed")
		return
	}

	rowCnt, err := result.RowsAffected()
	if err != nil {
		logger.WithField("error in rows affected", err.Error()).Error("Update DB Failed")
		return
	}
	logger.WithField("affected rows", rowCnt).Info("Update DB Successful")
	return
}

func updateQuery(resp model.XEcurrency) (string, []interface{}) {

	values := make([]string, 0, len(resp.To))
	args := make([]interface{}, 0, len(resp.To)*5)

	for _, to := range resp.To {
		values = append(values, "(?, ?, ?, ?, ?)")
		args = append(args, resp.From, to.Quotecurrency, to.Mid, resp.Timestamp, resp.Timestamp)
	}

	stmt := fmt.Sprintf(`INSERT INTO exchange_rates
			(from_currency,to_currency,rate,created_at,updated_at)
			VALUES %s`,
		strings.Join(values, ","))

	stmt = replaceSQL(stmt, "(?, ?, ?, ?, ?)", len(values))

	stmt += `ON CONFLICT (from_currency,to_currency) DO UPDATE
			SET rate=excluded.rate, updated_at = excluded.updated_at
			WHERE exchange_rates.from_currency = excluded.from_currency AND exchange_rates.to_currency = excluded.to_currency`
	return stmt, args
}

func replaceSQL(stmt, pattern string, len int) string {
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
