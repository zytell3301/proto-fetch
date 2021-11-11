package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v40/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"io/fs"
	"io/ioutil"
)

type Config struct {
	baseurl    string
	owner      string
	repository string
	token      string
	protoFiles []string
	output     string
}

func main() {
	configs, err := loadConfigs()
	ctx := context.Background()

	switch err != nil {
	case true:
		fmt.Printf("An error occurred while reading configs. Error: %v\n", err)
		return
	}

	var client *github.Client

	switch configs.token == "" {
	case true:
		client = github.NewClient(nil)
		break
	default:
		src := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: configs.token,
		})
		client = github.NewClient(oauth2.NewClient(ctx, src))
	}

	for _, file := range configs.protoFiles {
		protofile, _, _, err := client.Repositories.GetContents(ctx, configs.owner, configs.repository, file, &github.RepositoryContentGetOptions{})
		switch err != nil {
		case true:
			fmt.Printf("An error occurred while fetching a proto file. Error %v\n", err)
			return
		}

		content, _ := protofile.GetContent()
		switch err != nil {
		case true:
			fmt.Printf("An error occurred while reading a proto file. Error: %v\n", err)
		}
		ioutil.WriteFile(configs.output+*protofile.Name, []byte(content), fs.ModePerm)
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
		baseurl:    cfg.GetString("base-url"),
		owner:      cfg.GetString("repository-owner"),
		repository: cfg.GetString("repository"),
		token:      cfg.GetString("auth-token"),
		protoFiles: cfg.GetStringSlice("files"),
		output:     cfg.GetString("output-dir"),
	}, nil
}
