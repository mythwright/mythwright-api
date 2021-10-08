package main

import (
	"context"
	"github.com/Zyian/gw2-droplake-api/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"time"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME/.droplake/")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("droplake")

	viper.SetDefault("db_uri", "mongodb://localhost")
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	srv := api.NewServer(ctx)

	go srv.ListenAndServe(ctx)
	logrus.Info("Listening on ", viper.GetString("addr"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	<-c
	srv.Shutdown(ctx)
	logrus.Info("Got Interrupt Signal, Exiting")
}
