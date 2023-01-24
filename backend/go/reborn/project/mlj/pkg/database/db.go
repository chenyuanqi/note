package database

import (
	"fmt"
	"mlj/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	// dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	// 构建 DSN 信息
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		config.Get("database.mysql.username"),
		config.Get("database.mysql.password"),
		config.Get("database.mysql.host"),
		config.Get("database.mysql.port"),
		config.Get("database.mysql.database"),
		config.Get("database.mysql.charset"),
	)
	fmt.Println(dsn)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
