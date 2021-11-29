package db

import (
	m "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	User                 string
	Password             string
	Net                  string
	Addr                 string
	DBName               string
	AllowNativePasswords bool
	Params               Param
}

type Param struct {
	ParseTime string
}

func NewDB(config Config) *gorm.DB {
	mysqlConfig := m.Config{
		User:                 config.User,
		Passwd:               config.Password,
		Net:                  config.Net,
		Addr:                 config.Addr,
		AllowNativePasswords: config.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.Params.ParseTime,
		},
	}
	mysqlDB, err := gorm.Open(mysql.Open(mysqlConfig.FormatDSN()), &gorm.Config{})

	if err != nil {
		return nil
	}

	return mysqlDB
}
