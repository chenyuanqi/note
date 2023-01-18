package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/skyfile/cache2go"
	"gopkg.in/yaml.v2"
)

var (
	DefaultTableName = "config"
	once             sync.Once
	cache            *CacheInstance
)

type RedisConf struct {
	HostName string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	PassWord string `yaml:"password"`
	DataBase int    `yaml:"database"`
}

type CacheInstance struct {
	Cache *cache2go.CacheTable
	Redis *redis.Client
}

/**
获取缓存实例
*/
func Cache(tableNames ...string) *CacheInstance {
	once.Do(func() {
		tableName := DefaultTableName
		if len(tableNames) > 0 {
			tableName = tableNames[0]
		}
		cache = &CacheInstance{
			Cache: cache2go.Cache(tableName),
			Redis: RedisClient(CCRedisConf(tableName)),
		}
	})
	return cache
}

/**
获取redis实例
*/
func RedisClient(config ...*RedisConf) *redis.Client {
	var rdCf *RedisConf
	if len(config) > 0 {
		rdCf = config[0]
	} else {
		rdCf = CCRedisConf("config")
	}
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rdCf.HostName, rdCf.Port),
		Password: rdCf.PassWord,
		DB:       rdCf.DataBase,
	})
}

/**
从CC中获取redis配置信息
*/
func CCRedisConf(keyName string) *RedisConf {
	var rdCf *RedisConf
	conf, err := CCDefault.FromNacosService("redis", "DEFAULT_GROUP")
	if err != nil {
		// Log().Fatal(fmt.Sprintf("redisConf get from naoce error: %s", err.Error()))
		fmt.Sprintf("redisConf get from naoce error: %s", err.Error())
	}
	var rdCfMap map[string]RedisConf
	if err := yaml.Unmarshal([]byte(conf), &rdCfMap); err != nil {
		// Log().Fatal(fmt.Sprintf("json ummarshal redisConf error: %s", err.Error()))
		fmt.Sprintf("json ummarshal redisConf error: %s", err.Error())
	}
	if cf, ok := rdCfMap[keyName]; !ok {
		panic(fmt.Sprintf("redisConf['config'] not exist"))
	} else {
		rdCf = &cf
	}
	return rdCf
}

func (c *CacheInstance) Get(key string, ttl time.Duration, action ...func(k ...string) interface{}) (string, error) {
	if len(action) > 0 {
		c.Cache.SetDataLoader(func(_key interface{}, _value ...interface{}) *cache2go.CacheItem {
			if key != _key.(string) {
				return nil
			}
			// 优先读取redis数据，如果redis没有数据，则通过提供的方法获取
			res, err := c.Redis.Get(key).Result()
			if err == redis.Nil { // key 不存在，获取数据然后更新到redis
				item := cache2go.NewCacheItem(_key, ttl, action[0](_key.(string)))
				itemData := item.Data()
				if itemData == nil {
					return nil
				}
				// 判断如果不为 string 则将数据转换为json字符串
				itemType := fmt.Sprintf("%T", itemData)
				if itemType != "string" {
					itemDataByte, err := json.Marshal(itemData)
					if err != nil {
						return nil
					}
					res = string(itemDataByte)
				} else {
					res = itemData.(string)
				}
				// 将数据设置到redis中
				if err := c.Redis.SetNX(key, res, ttl).Err(); err != nil {
					return nil
				}
			} else if err != nil {
				return nil
			}
			return cache2go.NewCacheItem(_key, ttl, res)
		})
	}
	res, err := c.Cache.Value(key)
	if err != nil {
		return "", err
	}
	if res.Data() == nil {
		return "", errors.New("no data")
	}
	return res.Data().(string), nil
}

func (c *CacheInstance) Set(key string, value string, ttl time.Duration) {
	c.Redis.Set(key, value, ttl)
	c.Cache.Add(key, ttl, value)
}
