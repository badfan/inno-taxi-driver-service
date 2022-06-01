package app

import (
	"errors"

	"github.com/spf13/viper"
)

type APIConfig struct {
	APIPort  string
	TokenTTL int
}

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	SSLMode    string
}

func NewAPIConfig() (*APIConfig, error) {
	res := &APIConfig{}
	viper.AutomaticEnv()

	if !viper.IsSet("APIPORT") {
		return nil, errors.New("env var APIPORT is empty")
	}
	res.APIPort = viper.GetString("APIPORT")

	if !viper.IsSet("TOKENTTL") {
		return nil, errors.New("env var TOKENTTL is empty")
	}
	res.TokenTTL = viper.GetInt("TOKENTTL")

	return res, nil
}

func NewDBConfig() (*DBConfig, error) {
	res := &DBConfig{}
	viper.AutomaticEnv()

	if !viper.IsSet("DBUSER") {
		return nil, errors.New("env var DBUSER is empty")
	}
	res.DBUser = viper.GetString("DBUSER")

	if !viper.IsSet("DBPASSWORD") {
		return nil, errors.New("env var DBPASSWORD is empty")
	}
	res.DBPassword = viper.GetString("DBPASSWORD")

	if !viper.IsSet("DBHOST") {
		return nil, errors.New("env var DBHOST is empty")
	}
	res.DBHost = viper.GetString("DBHOST")

	if !viper.IsSet("DBPORT") {
		return nil, errors.New("env var DBPORT is empty")
	}
	res.DBPort = viper.GetString("DBPORT")

	if !viper.IsSet("DBNAME") {
		return nil, errors.New("env var DBNAME is empty")
	}
	res.DBName = viper.GetString("DBNAME")

	if !viper.IsSet("SSLMODE") {
		return nil, errors.New("env var SSLMODE is empty")
	}
	res.SSLMode = viper.GetString("SSLMODE")

	return res, nil
}
