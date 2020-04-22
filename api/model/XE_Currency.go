package model

import "time"

type XE_Currency_Response struct {
	Terms     string    `json:"terms"`
	Privacy   string    `json:"privacy"`
	From      string    `json:"from"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
	To        []struct {
		Quotecurrency string  `json:"quotecurrency"`
		Mid           float64 `json:"mid"1`
	} `json:"to"`
}
