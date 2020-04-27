package config

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//InitViper function to initialize viper
func InitConfig(path string) error {

	viper.SetConfigName(path)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		logger.WithField("error config file", err.Error()).Error("Cannot initialize config")
		return err
	}
	logger.WithField("msg", "initialize successfully").Info("Initialized config")
	return nil
}

//GetConfig method to get configs from config file
func GetConfig(keyName string) string {
	keyValue := viper.GetString(keyName)
	return keyValue
}

func GetStringSlice(keyName string) []string {
	keyValue := viper.GetStringSlice(keyName)
	return keyValue
}
