package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

var CONFIGURATION *Configuration

type Configuration struct {
	AppEnv      string `env:"APP_ENV"`
	LogLevel    string `env:"LOG_LEVEL"`
	RunningMode string `env:"RUNNING_MODE"` //fallback or main
}

func SetEnvConfig() {
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	logrus.Info("\nload configuration OK")
	CONFIGURATION = &cfg
}
