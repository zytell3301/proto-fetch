package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v40/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type Config struct {
	baseurl    string
	owner      string
	repository string
	token      string
}

func main() {
	configs, err := loadConfigs()
	ctx := context.Background()

	switch err != nil {
	case true:
		fmt.Printf("An error occurred while reading configs. Error: %v\n", err)
		return
	}

	src := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: configs.token,
	})
	httpClient := oauth2.NewClient(ctx, src)
	client := github.NewClient(httpClient)
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
