package db

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/domain/model"
	m "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(config *DB) (*gorm.DB, error) {
	mysqlConfig := m.Config{
		User:                 config.User,
		Passwd:               config.Password,
		Net:                  config.Net,
		Addr:                 config.Addr,
		DBName:               config.DBName,
		AllowNativePasswords: config.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.Params.ParseTime,
		},
	}

	mysqlDB, err := gorm.Open(mysql.Open(mysqlConfig.FormatDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return mysqlDB, nil
}

// AutoMigrate migrates the models
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		model.Product{},
		model.Link{},
		model.Order{},
		model.OrderItem{},
	)
}
