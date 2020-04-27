package service

import (
	"io/ioutil"
	"net/http"
	"testing"
	"xe-currency/config"

	"github.com/nbio/st"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/h2non/gock.v1"
)


func Test_Init(t *testing.T) {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})
	err := config.InitConfig("../config")
	if err != nil {
		t.Errorf("InitJob = %d; want ", err)
		return
	}
}


func TestAPIWithURL(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	response := `{"terms":"http://www.xe.com/legal/dfs.php","privacy":"http://www.xe.com/privacy.php","to":"CAD","amount":1.0,"timestamp":"2020-04-27T00:00:00Z","from":[{"quotecurrency":"USD","mid":0.7086741418}]}`

	gock.New("https://xecdapi.xe.com/v1/convert_from.json/").
		MatchHeader("Authorization", "Basic eHh4eHh4Njc2OTMwMDMyOjY4a3VuMm44OWhvZDB1ZXBzNGwyam4yYzMx").
		MatchParams(map[string]string{
			"from": "USD",
			"to":"CAD",
		}).
		Reply(200).
		JSON([]byte(response))

	req, err := http.NewRequest("GET", "https://xecdapi.xe.com/v1/convert_from.json/?from=USD&to=CAD", nil)
  	req.Header.Set("Authorization", "Basic eHh4eHh4Njc2OTMwMDMyOjY4a3VuMm44OWhvZDB1ZXBzNGwyam4yYzMx")

  	res, err := (&http.Client{}).Do(req)

	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)

	body, _ := ioutil.ReadAll(res.Body)
	st.Expect(t, string(body), response)

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)
}


func TestAPI(t *testing.T) {
	defer gock.Off()

	response := `{"terms":"http://www.xe.com/legal/dfs.php","privacy":"http://www.xe.com/privacy.php","to":"CAD","amount":1.0,"timestamp":"2020-04-27T00:00:00Z","from":[{"quotecurrency":"USD","mid":0.7086741418}]}`


	gock.New("https://xecdapi.xe.com/v1/convert_from.json/").
		MatchHeader("Authorization", "Basic eHh4eHh4Njc2OTMwMDMyOjY4a3VuMm44OWhvZDB1ZXBzNGwyam4yYzMx").
		MatchParams(map[string]string{
			"from": "USD",
			"to":"CAD",
		}).
		Reply(200).
		JSON([]byte(response))

	res,err := httpReqToXE("CAD","USD")
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)

	body, _ := ioutil.ReadAll(res.Body)
	st.Expect(t, string(body), response)

	st.Expect(t, gock.IsDone(), true)
}

func TestAPI_InvalidCurrency(t *testing.T) {
	defer gock.Off()

	response := `{"code":7,"message":"No CADK found on 2020-04-27T00:00:00Z","documentation_url":"https://xecdapi.xe.com/docs/v1/"}`


	gock.New("https://xecdapi.xe.com/v1/convert_from.json/").
		MatchHeader("Authorization", "Basic eHh4eHh4Njc2OTMwMDMyOjY4a3VuMm44OWhvZDB1ZXBzNGwyam4yYzMx").
		MatchParams(map[string]string{
			"from": "USD",
			"to":"CADK",
		}).
		Reply(200).
		JSON([]byte(response))

	res,err := httpReqToXE("CADK","USD")
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)

	body, _ := ioutil.ReadAll(res.Body)
	st.Expect(t, string(body), response)

	st.Expect(t, gock.IsDone(), true)
}

func TestAPI_Authorization(t *testing.T) {
	defer gock.Off()

	response := `{"code":1, "message": "Bad credentials", "documentation_url": "https://xecdapi.xe.com/docs/v1/" } `


	gock.New("https://xecdapi.xe.com/v1/convert_from.json/").
		MatchHeader("Authorization", "Basic eHh4eHh4Njc2OTMwMDMyOjY4a3VuMm44OWhvZDB1ZXBzNGwyam4yYzMx").
		MatchParams(map[string]string{
			"from": "USD",
			"to":"CAD",
		}).
		Reply(401).
		JSON([]byte(response))

	res,err := httpReqToXE("CAD","USD")
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 401)

	body, _ := ioutil.ReadAll(res.Body)
	st.Expect(t, string(body), response)

	st.Expect(t, gock.IsDone(), true)
}