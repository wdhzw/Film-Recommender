package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	globalConfig *Config

	NilConfigError = fmt.Errorf("nil config")
)

type DynamoDBConfig struct {
	Region            string `yaml:"Region"`
	AccessID          string `yaml:"AccessID"`
	AccessSecret      string `yaml:"AccessSecret"`
	ConnectionTimeout string `yaml:"ConnectionTimeout,omitempty"`
}

type CognitoConfig struct {
	Region      string `yaml:"Region"`
	AppClientID string `yaml:"AppClientID"`
}

type Config struct {
	Dynamo  DynamoDBConfig `yaml:"Dynamo"`
	Cognito CognitoConfig  `yaml:"Cognito"`
}

func Init() {
	globalConfig = loadConfig()
}

func GetConfig() *Config {
	return globalConfig
}

func loadConfig() *Config {
	filepath := getConfigFilepath()
	log.Printf("config filepath:%s\n", filepath)
	f, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	var config Config
	if err = yaml.Unmarshal(f, &config); err != nil {
		panic(err)
	}
	return &config
}

func getConfigFilepath() string {
	return "conf/conf.yml"
}
