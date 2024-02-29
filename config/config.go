package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

var (
	ErrConfigPathNotSet       = errors.New("CONFIG_PATH is not set")
	ErrConfigFileNotExists    = errors.New("config file does not exist")
	ErrCouldNotReadConfigFile = errors.New("could not read config file")
)

func Init[T any]() (*T, error) {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		return nil, ErrConfigPathNotSet
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, ErrConfigFileNotExists
	}

	cfg := new(T)

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, ErrCouldNotReadConfigFile
	}

	return cfg, nil
}

func MustInit[T any]() *T {
	cfg, err := Init[T]()

	if err != nil {
		panic(err)
	}

	return cfg
}
