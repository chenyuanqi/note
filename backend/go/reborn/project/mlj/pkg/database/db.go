package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB, _ = GetDB()

func GetDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
