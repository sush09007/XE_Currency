package main

import (
	"xe-currency/config"
	"xe-currency/service"

	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	err := config.InitConfig("config")
	if err != nil {
		logger.WithField("error in config file", err.Error()).Error("Exit")
		return
	}

	service.InitJob()

}
