package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/leapsquare/sheet-service/config"
	"log"
	"os"
)

func LoadAndParseCfgFile() (*config.Config, error) {
	configPath := GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
		return nil, err
	}
	return config.ParseConfig(cfgFile)
}

func GetJsonValidator() *validator.Validate {
	return validator.New()
}
