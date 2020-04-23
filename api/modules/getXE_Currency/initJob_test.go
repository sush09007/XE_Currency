package getXE_Currency

import (
	"XE_Currency/api/utils"
	"testing"
)

func TestInitJob(t *testing.T) {
	utils.InitViper()
	got := InitJob()
	if got != 1 {
		t.Errorf("InitJob = %d; want ", got)
	}
}
