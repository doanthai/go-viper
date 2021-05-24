package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig(env string)  {
	switch env {
	case "test", "local", "stg", "prod":
		viper.SetConfigName(env)
	default:
		panic(fmt.Sprint("Env name incorrect"))
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
