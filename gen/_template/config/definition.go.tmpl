package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	App   `json:"app"`
	Redis `json:"redis"`
}

type App struct {
	Port uint `json:"port"`
}

type Redis struct {
	Url string `yaml:"url"`
}

func Load() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("./config")

	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	var conf *Config

	if err := v.Unmarshal(&conf); err != nil {
		log.Fatalln(err)
	}

	return conf
}
