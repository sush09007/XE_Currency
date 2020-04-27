package main

import (
	"testing"
	"xe-currency/config"
	"xe-currency/service"

	logger "github.com/sirupsen/logrus"
)

func Test_Init(t *testing.T) {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})
	err := config.InitConfig()
	if err != nil {
		t.Errorf("InitJob = %d; want ", err)
		return
	}
}
func Test_Job(t *testing.T) {
	got := service.InitJob()
	if got != 1 {
		t.Errorf("InitJob = %d; want ", got)
	}
}

// func Benchmark_InitJob(b *testing.B) {
//         for n := 0; n < b.N; n++ {
//                 service.InitJob()
//         }
// }
