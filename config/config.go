package config

import "github.com/spf13/viper"

type Config struct {
	Host       string `mapstructure:"HOST"`
	Port       string `mapstructure:"PORT"`
	DbHost     string `mapstructure:"DBHOST"`
	DbUser     string `mapstructure:"DBUSER"`
	DbName     string `mapstructure:"DBNAME"`
	DbPassword string `mapstructure:"DBPASSWORD"`
	DbPort     string `mapstructure:"DBPORT"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
