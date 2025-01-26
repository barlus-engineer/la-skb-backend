package config

import (
	"laskb-server-api/pkg/logger"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"postgres"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}

func LoadConfig() *Config {
	var (
		cfg Config
		filePath = "config.yml"
	)
	file, err := os.Open(filePath)
	if err != nil {
		logger.Fatalf("Failed to open config file '%s'. Error: %v", filePath, err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		logger.Fatalf("Failed to decode YAML file. Error: %v", err)
	}

	return &cfg
}