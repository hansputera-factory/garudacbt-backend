package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server
		Database *Database
		Secrets  *Secrets
		Mode     string
	}

	Server struct {
		Host string
		Port int
	}

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}

	Secrets struct {
		JwtKey       string
		AuthorizeKey string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AddConfigPath("/etc/garuda-cbt")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
