package db

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB
	Redis
}

type DB struct {
	User                 string
	Password             string
	Net                  string
	Addr                 string
	DBName               string
	AllowNativePasswords bool
	Params               Param
}

type Redis struct {
	Port string
}

type Param struct {
	ParseTime string
}

func NewConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: DB{
			User:                 viper.GetString("db.user"),
			Password:             viper.GetString("db.password"),
			Net:                  viper.GetString("db.net"),
			Addr:                 viper.GetString("db.addr"),
			DBName:               viper.GetString("db.name"),
			AllowNativePasswords: viper.GetBool("db.allowNativePasswords"),
			Params: Param{
				ParseTime: viper.GetString("db.params.parseTime"),
			},
		},
		Redis: Redis{
			Port: viper.GetString("redis.port"),
		},
	}, nil

}
