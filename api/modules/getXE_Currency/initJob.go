package getXE_Currency

import (
	"log"
	"strings"
	"sync"
	"time"
	"XE_Currency/api/utils"
)

var Currencies = utils.GetStringSlice("currency")

func InitJob() int {
	start := time.Now()

	var w sync.WaitGroup
	for i, value := range Currencies {
		from := value
		to := strings.Join(Currencies, ",")
		w.Add(1)
		go RequestXEForData(from, to, &w, i)
	}
	w.Wait()

	elapsed := time.Now().Sub(start)
	log.Print("Total time taken: ", elapsed)
	return 1
}
