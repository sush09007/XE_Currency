package service

import (
	"net/http"
	"xe-currency/config"

	logger "github.com/sirupsen/logrus"
)

func httpReqToXE(to, from string) (resp *http.Response, err error) {

	url := config.GetConfig("xe_account.xe_url")
	username := config.GetConfig("xe_account.xe_account_id")
	passwd := config.GetConfig("xe_account.xe_account_key")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("to", to)
	q.Add("from", from)

	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(username, passwd)

	// logger.WithField("req to api",req).Info("API Failed")
	resp, err = client.Do(req)

	if err != nil {
		logger.WithField("error from api", err.Error()).Error("API Failed")
		return
	}
	return
}
