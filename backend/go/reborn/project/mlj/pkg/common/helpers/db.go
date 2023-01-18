package helpers

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

	"mlj/pkg/common/consts"
)

type DBOpt struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
	Charset  string
}

func NewDB(opt *DBOpt) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.DBName,
		opt.Charset,
	)
	log.Printf("db connect, dsn: %s\n", dsn)

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &logger.Recorder,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err))
	}
	if consts.EnvModeIsDev {
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
	return db
}

type RedisOpt struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewRedis(opt *RedisOpt) *redis.Client {
	addr := fmt.Sprintf("%s:%d",
		opt.Host,
		opt.Port,
	)
	log.Printf("redis connect, dsn=%s/%d\n", addr, opt.DB)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: opt.Password, // no password set
		DB:       opt.DB,       // use default DB
	})
	pong, err := redisClient.Ping().Result()
	log.Println("redis ping:", pong)
	if err != nil && err != redis.Nil {
		panic(fmt.Sprintf("redis err: %s", err))
	}
	return redisClient
}
