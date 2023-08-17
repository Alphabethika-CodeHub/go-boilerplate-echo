package utilities

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadViper() {
	root, _ := os.Getwd()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(root)
	viper.SetConfigName("app.dev")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file : ", err.Error())
		os.Exit(0)
	}
}
