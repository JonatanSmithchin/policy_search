package Mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type DBConfig struct {
	DriverName string
	Host       string
	UserName   string
	Password   string
	Database   string
	Charset    string
}

var DB *gorm.DB

func init() {
	tempDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:44913730@tcp(121.37.119.47:3306)/db_policy_search?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{})

	if err != nil {
		log.Panicf("cannot open mysql: %v", err)
	}
	sqlDB, err := tempDb.DB()
	if err != nil {
		log.Printf("database setup error: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = tempDb
}
