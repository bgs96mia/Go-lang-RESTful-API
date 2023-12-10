package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.AddConfigPath(".")
	config.SetConfigName("config")
	config.SetConfigType("json")

	err := config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return config
}
