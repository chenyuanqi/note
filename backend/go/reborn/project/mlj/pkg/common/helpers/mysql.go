package helpers

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var mysqlDBs sync.Map

type originConf struct {
	Dsn         string
	Username    string
	Password    string
	Charset     string
	TablePrefix string `yaml:"tablePrefix"`
}

func GetMysqlDB(dbKey string) *gorm.DB {
	if db, exist := mysqlDBs.Load(dbKey); exist {
		return db.(*gorm.DB)
	}
	var originConf map[string]originConf
	if err := CCDefault.Get("mysql", "", &originConf); err != nil {
		// Log().Fatal(fmt.Sprintf("mysql config from CC ERROR: %s", err.Error()))
		fmt.Sprintf("mysql config from CC ERROR: %s", err.Error())
	}
	cf, ok := originConf[dbKey]
	if !ok {
		// Log().Fatal(fmt.Sprintf("originConf['%s'] not exist", dbKey), originConf)
		fmt.Sprintf("originConf['%s'] not exist", dbKey)
	}
	arr := strings.Split(cf.Dsn, ";")
	host := strings.Split(arr[0], "=")
	dbName := strings.Split(arr[1], "=")
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true&loc=Asia%%2FShanghai", cf.Username, cf.Password, host[1], dbName[1], cf.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &logger.Recorder,
	})
	if err != nil {
		// Log().Fatal(fmt.Sprintf("Mysql Instance Failed: %v", err.Error()))
		fmt.Sprintf("Mysql Instance Failed: %v", err.Error())
		//not necessary
		return nil
	}
	if os.Getenv("DB_DEBUG") != "" {
		db = db.Debug()
	}
	db = db.Set("gorm:association_autocreate", false)
	db = db.Set("gorm:association_autoupdate", false)
	db = db.Set("gorm:association_save_reference", false)
	db = db.Omit(clause.Associations)
	db = db.Session(&gorm.Session{
		NewDB:                true,
		FullSaveAssociations: false,
	})

	mysqlDBs.Store(dbKey, db)
	return db
}
