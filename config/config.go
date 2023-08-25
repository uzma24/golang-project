package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURI string `mapstructure:"DEV_DATABASE_URI"`
	RedisURI    string `mapstructure:"REDIS_URI"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("error in reading config:::", err)
		return
	}
	err = viper.Unmarshal(&config)
	fmt.Println(viper.Unmarshal(&config))

	return
}
