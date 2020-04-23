package getXE_Currency

import (
	"XE_Currency/api/utils"
	"log"
	"strings"
	"sync"
	"time"
)

var Currencies []string

func InitCurrencies() {
	Currencies = utils.GetStringSlice("currency")
}

type Result struct {
	message string
	err     error
}

func InitJob() int {
	start := time.Now()
	// c := make(chan Result)

	log.Print("Currencies", Currencies)
	var w sync.WaitGroup
	for i, value := range Currencies {
		from := value
		to := strings.Join(Currencies, ",")
		w.Add(1)
		go RequestXEForData(from, to, &w, i)
	}
	// i, ok := <-c
	// log.Printf("OK",ok,i)
	// close(c)
	// for v := range c {
	// 	log.Print(v.message,v.err)
	// 	if v.err != nil {
	// 		log.Printf("read value %s from channel", v)
	// 		close(c)
	// 		return -1
	// 	}
	// }

	// log.Println(len(c))
	w.Wait()

	elapsed := time.Now().Sub(start)
	log.Print("Total time taken: ", elapsed)
	return 1
}
