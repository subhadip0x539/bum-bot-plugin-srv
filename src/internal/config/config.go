package config

import (
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Mongo struct {
	URI      string
	Database string
	Timeout  time.Duration
}

type Server struct {
	Host           string
	Port           int
	TrustedProxies []string
}

type Config struct {
	Server
	Mongo
}

func (c *Config) Register(configFile string, configType string, mode string) error {
	baseDir, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.AddConfigPath(baseDir)
	viper.SetConfigName(configFile)
	viper.SetConfigType(configType)

	if mode == "release" {
		viper.AutomaticEnv()
	} else {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	c.Mongo = Mongo{
		URI:      viper.GetString("MONGO_URI"),
		Database: viper.GetString("MONGO_DATABASE"),
		Timeout:  viper.GetDuration("MONGO_TIMEOUT"),
	}
	c.Server = Server{
		Host:           viper.GetString("SERVER_HOST"),
		Port:           viper.GetInt("SERVER_PORT"),
		TrustedProxies: strings.Split(viper.GetString("SERVER_TRUSTED_PROXIES"), ","),
	}

	return nil
}

func (c Config) Get() Config {
	config := c

	return config
}

func NewConfig() *Config {
	return &Config{}
}
