package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"xe-currency/config"
	"xe-currency/model"

	logger "github.com/sirupsen/logrus"
)

func httpReqToXE(to, from string) (xe_resp model.XEcurrency, err error) {

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

	resp, err := client.Do(req)

	if err != nil {
		logger.WithField("error from api", err.Error()).Error("API Failed")
		return
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.WithField("response", err.Error()).Error("Reading Response Failed")
		return
	}

	err = json.Unmarshal(r, &xe_resp)
	if err != nil {
		logger.WithField("error in unmarshalling response", err.Error()).Error("Reading Response Failed")
		return
	}

	if xe_resp.From == "" || len(xe_resp.To) == 0 {
		err_resp := model.ErrorResponse{}
		err = json.Unmarshal(r, &err_resp)
		if err != nil {
			logger.WithField("error in unmarshalling response", err.Error()).Error("Reading Response Failed")
			return
		}
		logger.WithField("error in in response from api", err_resp.Message).Error("API Failed")
		return
	}
	return
}
