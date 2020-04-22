package getXE_Currency

import (
	"testing"
	"XE_Currency/api/utils"
)


func TestInitJob(t *testing.T) {
	utils.InitViper()
	Currencies = []string{"AED", "CUP", "AFN"}
	got := InitJob()
	if got != 1 {
        t.Errorf("InitJob = %d; want ", got)
    }
}
