package getXE_Currency

import "testing"

func TestInitJob(t *testing.T) {
	got := InitJob()
	if got != "job executed successfully" {
        t.Errorf("InitJob = %d; want ", got)
    }
}
