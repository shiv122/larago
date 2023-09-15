package config

import "github.com/spf13/viper"

var Port string
var DatabaseConfig DatabaseConfiguration

type DatabaseConfiguration struct {
	Database string
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadConfig() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Port = viper.GetString("port")
	DatabaseConfig = DatabaseConfiguration{
		Database: viper.GetString("database.database"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		DBName:   viper.GetString("database.dbname"),
	}
}
