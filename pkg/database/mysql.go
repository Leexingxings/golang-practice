package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"practice/pkg/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var database *gorm.DB
	var err error

	dsn := "root:lijiaxing@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Panic("Failed to connect database")
	}
	//sqlDB, err := database.DB()
	//if err != nil {
	//	logger.Panic("Failed to get database instance")
	//}
	//defer sqlDB.Close()
	//
	//// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//sqlDB.SetMaxIdleConns(10)
	//
	//// SetMaxOpenConns sets the maximum number of open connections to the database.
	//sqlDB.SetMaxOpenConns(100)
	//
	//// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//sqlDB.SetConnMaxLifetime(time.Hour)

	DB = database
}
