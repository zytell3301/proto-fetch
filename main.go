package main

import "github.com/spf13/viper"

type Config struct {
	repository string
}

func main() {

}

func loadConfigs() (Config,error){
	cfg := viper.New()
	cfg.SetConfigName("proto-fetch")
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath("./")
	err := cfg.ReadInConfig()
	switch err != nil {
	case true:
		return Config{},err
	}

	return Config{
		repository: cfg.GetString("repository"),
	}, nil
}