package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v40/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"io/fs"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	baseurl                string
	owner                  string
	repository             string
	token                  string
	protoFiles             []string
	output                 string
	afterFetchCommands     []string
	beforeFetchCommands    []string
	environmentalVariables []string
}

type File struct {
	name        string
	destination string
}

var environmentalVariables map[string]string

func main() {
	configs, err := loadConfigs()
	ctx := context.Background()
	parseVariables(&configs)
	configs.output = fillVariablePlaceHolders(configs.output)

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

	executeCommands(configs.beforeFetchCommands)

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
		err = os.MkdirAll(configs.output, fs.ModePerm)
		os.WriteFile(configs.output+*protofile.Name, []byte(content), fs.ModePerm)
		switch err != nil {
		case true:
			fmt.Printf("An error occurred while creating files. Error: %v\n", err)
		}
	}

	executeCommands(configs.afterFetchCommands)
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
		baseurl:                cfg.GetString("base-url"),
		owner:                  cfg.GetString("repository-owner"),
		repository:             cfg.GetString("repository"),
		token:                  cfg.GetString("auth-token"),
		protoFiles:             cfg.GetStringSlice("files"),
		output:                 cfg.GetString("output-dir"),
		beforeFetchCommands:    cfg.GetStringSlice("before-fetch-commands"),
		afterFetchCommands:     cfg.GetStringSlice("after-fetch-commands"),
		environmentalVariables: cfg.GetStringSlice("env-variables"),
	}, nil
}

func executeCommands(commands []string) {
	for _, cmd := range commands {
		cmd = fillVariablePlaceHolders(cmd)
		command := strings.SplitN(cmd, " ", 2)
		out := make([]byte, 0)
		switch len(command) == 1 {
		case true:
			out, _ = exec.Command(command[0]).Output()
			break
		default:
			out, _ = exec.Command(command[0], command[1]).Output()
		}
		fmt.Println(string(out))
	}
}

func parseVariables(configs *Config) {
	variables := make(map[string]string)
	for _, variable := range configs.environmentalVariables {
		keyPair := strings.SplitN(variable, "=", 2)
		variables[keyPair[0]] = keyPair[1]
	}

	environmentalVariables = variables
}

func fillVariablePlaceHolders(expression string) string {
	for key, value := range environmentalVariables {
		expression = strings.ReplaceAll(expression, key, value)
	}
	return expression
}
