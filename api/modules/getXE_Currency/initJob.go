package getXE_Currency

import (
	"log"
	"strings"
	"sync"
	"time"
)

var Currencies = []string{"AED", "CUP", "AFN", "ETB", "ALL", "AMD", "AOA", "ARS", "AZN", "BAM", "BBD", "BDT", "BGN", "IQD", "BMD", "IRR", "BIF", "BRL", "BSD", "BTN", "BYN", "CAD", "BZD", "KPW", "JOD", "COP", "CRC", "CVE", "CZK", "DOP", "DZD", "EGP", "GBP", "GEL", "AWG", "GHS", "GIP", "GTQ", "GYD", "HKD", "HNL", "HRK", "HUF", "CUC", "ILS", "IMP", "INR", "BOB", "JEP", "JMD", "KES", "KGS", "FKP", "CHF", "ERN", "GGP", "BND", "CDF", "IDR", "CLP", "GNF", "JPY", "KMF", "SPL", "PYG", "TZS", "MRU", "KYD", "KZT", "MDL", "LKR", "LRD", "LSL", "RUB", "MGA", "SHP", "MMK", "MNT", "MOP", "MUR", "MVR", "MWK", "MXN", "MYR", "NAD", "NGN", "NIO", "NPR", "NZD", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "KWD", "RON", "RSD", "SYP", "LYD", "SAR", "SBD", "SDG", "SEK", "SGD", "TWD", "SOS", "SRD", "SZL", "TJS", "TMT", "STN", "TOP", "TRY", "TVD", "TND", "MKD", "UAH", "UGX", "UYU", "UZS", "USD", "LAK", "RWF", "KRW", "BHD", "OMR", "BWP", "XCD", "CNY", "YER", "ZAR", "ZMW", "ANG", "FJD", "GMD", "HTG", "KHR", "LBP", "MAD", "MZN", "QAR", "SCR", "SLL", "THB", "TTD", "AUD", "DKK", "NOK", "SVC", "VEF", "WST", "ZWD", "EUR", "VES", "XOF", "XPF", "DJF", "ISK", "VUV", "XAF", "VND"}

func InitJob() {
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
}
