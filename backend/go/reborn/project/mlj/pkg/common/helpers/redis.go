package helpers

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	instanceCache map[string]*redis.Client
	rwLock        sync.RWMutex
)

type redisConf struct {
	Host     string `yaml:"hostname"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

func init() {
	instanceCache = make(map[string]*redis.Client)
}

func GetRedis(dbKey string) *redis.Client {
	rwLock.RLock()
	instance, ok := instanceCache[dbKey]
	rwLock.RUnlock()
	if ok {
		return instance
	}
	rwLock.Lock()
	defer rwLock.Unlock()

	allConfig := make(map[string]redisConf)
	if err := CCDefault.Get("redis", "", &allConfig); err != nil {
		// Log().Fatal(fmt.Sprintf("redis config from CC ERROR: %s", err.Error()))
		fmt.Sprintf("redis config from CC ERROR: %s", err.Error())
	}

	var (
		redisConfig redisConf
	)
	if redisConfig, ok = allConfig[dbKey]; !ok {
		// Log().Fatal(fmt.Sprintf("originConf['%s'] not exist", dbKey), redisConfig)
		fmt.Sprintf("originConf['%s'] not exist", dbKey)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	})

	instanceCache[dbKey] = rdb

	return rdb
}
