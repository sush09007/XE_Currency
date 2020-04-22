package main

import (
	"testing"
	"XE_Currency/api/utils"
		"XE_Currency/api/modules/getXE_Currency"
)


func Test_main(t *testing.T) {
	utils.InitViper()
	getXE_Currency.Currencies = []string{"AED", "CUP", "AFN"}
	got := getXE_Currency.InitJob()
	if got != 1 {
        t.Errorf("InitJob = %d; want ", got)
    }
}
