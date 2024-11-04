package main

import (
	"errors"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	drivers "github.com/leo-andrei/telemetry/drivers"
	logger "github.com/leo-andrei/telemetry/log"
)

type Env struct {
	LogLevel   string `env:"LOG_LEVEL"`
	FilePath   string `env:"FILEPATH"`
	DriverType string `env:"DRIVER_TYPE" envDefault:"cli"`
}

func Load() (Env, error) {
	cfg := Env{}
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	if err := env.Parse(&cfg); err != nil {
		return Env{}, err
	}
	return cfg, nil
}

func InitLogger() (*logger.Logger, error) {
	envLoaded, err := Load()
	if err != nil {
		return nil, err
	}

	var driver logger.Driver
	switch envLoaded.DriverType {
	case "cli":
		driver = drivers.CLIDriver{}
	case "file":
		driver = drivers.FileDriver{FilePath: envLoaded.FilePath}
	case "json":
		driver = drivers.JSONFileDriver{FilePath: envLoaded.FilePath}
	default:
		return nil, errors.New("Unknown driver type")
	}

	return logger.NewLogger(driver), nil
}

func main() {
	loggerService, err := InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	loggerService.Log(logger.INFO, "This is a info message", map[string]interface{}{"CustomerId": "12345", "Environment": "production"})
}
