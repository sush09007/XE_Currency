package main

import (
	"XE_Currency/api/modules/getXE_Currency"
	"XE_Currency/api/utils"
	"testing"
)

func Test_InitJob(t *testing.T) {
	utils.InitViper()
	getXE_Currency.Currencies = []string{"AED", "CUP", "AFN"}
	got := getXE_Currency.InitJob()
	if got != 1 {
		t.Errorf("InitJob = %d; want ", got)
	}
}

// func Benchmark_InitJob(b *testing.B) {
// 	utils.InitViper()
// 	getXE_Currency.Currencies = []string{"AED", "CUP", "AFN", "ETB", "ALL", "AMD", "AOA", "ARS", "AZN", "BAM"}
//         for n := 0; n < b.N; n++ {
//                 getXE_Currency.InitJob()
//         }
// }
