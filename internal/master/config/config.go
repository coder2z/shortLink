package config

import (
	"fmt"
	"github.com/spf13/viper"
	"shortLink/internal/master/config/server"
	"shortLink/pkg/client/database"
	"shortLink/pkg/client/redis"
	"shortLink/pkg/log"
)

const (
	// DefaultConfigurationName is the default name of configuration
	defaultConfigurationName = "config"

	// DefaultConfigurationPath the default location of the configuration file
	defaultConfigurationPath = "./config"
)

type Cfg struct {
	Server *server.Options
	Log    *log.Options
	Redis  *redis.Options
	DB     *database.Options
}

func cfg() *Cfg {
	return &Cfg{
		Server: server.NewServerOptions(),
		Log:    log.NewLogOptions(),
		Redis:  redis.NewRedisOptions(),
		DB:     database.NewDatabaseOptions(),
	}
}

func TryLoadFromDisk() (*Cfg, error) {
	viper.SetConfigName(defaultConfigurationName)
	viper.AddConfigPath(defaultConfigurationPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err
		} else {
			return nil, fmt.Errorf("error parsing configuration file %s", err)
		}
	}

	conf := cfg()

	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
