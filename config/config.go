package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server Server `yaml:"server"`
	}

	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"post"`
	}
)

func newConfig() (*Config, error) {
	configViper := viper.New()

	// settings
	configViper.SetConfigName("config")
	configViper.SetConfigType("yaml")
	configViper.AddConfigPath("./config/")

	return ParseConfig(configViper)
}

func ParseConfig(configViper *viper.Viper) (*Config, error) {
	// read
	if err := configViper.ReadInConfig(); err != nil {
		return nil, err
	}

	// parse
	parsedConfig, err := ParseYaml(configViper)
	if err != nil {
		return nil, err
	}
	return parsedConfig, nil
}

func ParseYaml(configViper *viper.Viper) (*Config, error) {
	config := &Config{}
	if err := configViper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}
