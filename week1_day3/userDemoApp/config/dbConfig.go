package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbName   string
	password string
}

func New() (*gorm.DB, error) {
	return gorm.Open("mysql", buildDSN())
}

func configureDb() *dbConfig {
	return &dbConfig{
		host:     "localhost",
		port:     3306,
		user:     "root",
		dbName:   "userDemoApp",
		password: "yashDBraj",
	}
}

func buildDSN() string {
	config := configureDb()
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.user,
		config.password,
		config.host,
		config.port,
		config.dbName,
	)
}
