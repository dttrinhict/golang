package configs

import (
	"github.com/spf13/viper"
)

var conf *Config

func LoadConfig(path string) (*Config, error) {
	if conf == nil {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")

		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			return conf, err
		}
		err = viper.Unmarshal(&conf)
		if err != nil {
			return conf, err
		}
		return conf, nil
	}
	return conf, nil
}