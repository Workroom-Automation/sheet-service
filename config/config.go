package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
)

// Config struct
type Config struct {
	Server   ServerConfig
	Logger   Logger
	Postgres Postgres
	Auth     Auth
}

// Auth config
type Auth struct {
	Auth0KeySetUrl string
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Postgres config
type Postgres struct {
	DSN           string
	MigrationPath string
}

// ServerConfig struct
type ServerConfig struct {
	Port              string
	CtxDefaultTimeout time.Duration
	Mode              string
	Debug             bool
}

// LoadConfig file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvPrefix("CF")
	v.SetEnvKeyReplacer(replacer)

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error default config file: %w ", err))
	}
	configFileName := os.Getenv("ENV")
	if configFileName == "" {
		configFileName = "config"
	}
	v.SetConfigName(configFileName)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err = v.MergeInConfig()
	if err != nil {
		fmt.Println("There are errors while reading config overrides " + err.Error())
	}
	return v, nil
}

// ParseConfig file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
