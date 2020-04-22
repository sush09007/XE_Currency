package getXE_Currency

import (
	"XE_Currency/api/model"
	"XE_Currency/api/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func RequestXEForData(from, to string, wg *sync.WaitGroup, i int) (XE_Currency_Response model.XE_Currency_Response, err error) {

	url := utils.GetConfig("xe_account.xe_url")
	username := utils.GetConfig("xe_account.xe_account_id")
	passwd := utils.GetConfig("xe_account.xe_account_key")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("to", to)
	q.Add("from", from)
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	responseByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error in reading Response", err)
		return
	}
	log.Print(string(responseByte))
	err = json.Unmarshal(responseByte, &XE_Currency_Response)
	if err != nil {
		log.Fatal("Error in UnMarshaling Response", err)
		return
	}
	err = updateDB(XE_Currency_Response)
	if err != nil {
		log.Fatal("Error in Insert", err)
		return
	}
	wg.Done()
	return
}
