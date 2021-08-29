package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME/.droplake/")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("droplake")
}

func main() {
	if ll := os.Getenv("LOG_LEVEL"); ll != "" {
		level, err := logrus.ParseLevel(ll)
		if err == nil {
			logrus.SetLevel(level)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("unable to read config file: (%v)", err)
		logrus.Debugf("config loading failed; using default values")
	}

}
