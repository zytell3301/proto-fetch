package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	repository string
	token      string
}

func main() {
	configs, err := loadConfigs()

	switch err != nil {
	case true:
		fmt.Printf("An error occurred while reading configs. Error: %v\n", err)
		return
	}
}

func loadConfigs() (Config, error) {
	cfg := viper.New()
	cfg.SetConfigName("proto-fetch")
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath("./")
	err := cfg.ReadInConfig()
	switch err != nil {
	case true:
		return Config{}, err
	}

	return Config{
		repository: cfg.GetString("repository"),
	}, nil
}
