package getXE_Currency

import (
	"XE_Currency/api/model"
	"XE_Currency/api/utils"
	"encoding/json"
	// "errors"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)
// ch chan Result
func RequestXEForData(from, to string, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	// result := Result{}

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
		log.Fatal("Error While getting data from API: %s", err.Error())
		// result.message = "Error While getting data from API"
		// result.err = err
		// ch <- result
		return

	}
	responseByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error in reading Response", err.Error())
		// result.message = "Error While getting data from API"
		// result.err = err
		// ch <- result
		return

	}
	XE_Currency_Response := model.XE_Currency_Response{}

	err = json.Unmarshal(responseByte, &XE_Currency_Response)
	if err != nil {
		log.Fatal("Error in UnMarshaling Response: %s", err.Error())
		// result.message = "Error While getting data from API"
		// result.err = err
		// ch <- result
		return

	}

	if XE_Currency_Response.From == "" || len(XE_Currency_Response.To) == 0 {
		XE_Currency_ErrorResponse := model.ErrorResponse{}
		err = json.Unmarshal(responseByte, &XE_Currency_ErrorResponse)
		if err != nil {
			log.Fatal("Error in UnMarshaling Response: %s", err.Error())
			// result.message = "Error While getting data from API"
			// result.err = err
			// ch <- result
			return
		}
		log.Printf("Error in Response from API : %s", XE_Currency_ErrorResponse.Message)
		// result.message = "Error While getting data from API"
		// result.err = errors.New(fmt.Sprintf("Error in Response from API : %s", XE_Currency_ErrorResponse.Message))
		// ch <- result
		return

	}
	err = updateDB(XE_Currency_Response)
	if err != nil {
		log.Fatal("Error in Insert", err)
		// result.message = "Error While getting data from API"
		// result.err = err
		// ch <- result
		return
	}
	log.Print("goroutine-->")
	log.Print(i)
	// if err == nil {
	// 	var e error
	// 	log.Println(e)
	// 	result.message = "OK"
	// 	result.err = e
	// 	ch <- result
	// }
	// wg.Done()
	return
}
