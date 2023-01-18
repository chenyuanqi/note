package CC

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"mlj/pkg/common/consts"
	"mlj/pkg/common/helpers"

	"github.com/skyfile/cache2go"
	"gopkg.in/yaml.v2"
)

const (
	ConfigServiceDev     = "http://config.maliujia"
	ConfigServiceRelease = "http://nacos-k8s.default"
	DefaultGroupID       = "DEFAULT_GROUP"
)

var instance *configCenter

type configCenter struct {
	Host  string
	Cache *cache2go.CacheTable
}

func init() {
	host := ConfigServiceDev
	if consts.EnvMode == consts.EnvModeProd {
		host = ConfigServiceRelease
	}
	instance = &configCenter{
		Host:  host,
		Cache: cache2go.Cache("confCenter"),
	}
}

func Get(dataId, group string, outStruct interface{}) error {
	if group == "" {
		group = DefaultGroupID
	}
	key := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s", dataId, group))))
	//尝试内存缓存中取值
	result, err := instance.Cache.Value(key)
	if err == nil {
		return yaml.Unmarshal([]byte(result.Data().(string)), outStruct)
	}
	//尝试CC服务器取值
	nacos, err := FromNacosService(dataId, group)
	if err == nil {
		if err := yaml.Unmarshal([]byte(nacos), outStruct); err != nil {
			return err
		}
		// 缓存到内存表中
		instance.Cache.Add(key, 60*time.Second, nacos)
		// 缓存到本地文件中
		if _, err = localCache(key, nacos); err == nil {
			return nil
		} else {
			fmt.Printf("cc write to localCache failed: %+v\n", err)
		}
	}

	//尝试本地文件读取
	local, err := localCache(key)
	if err != nil {
		fmt.Printf("try read localCache err: %+v", err)
	} else {
		return yaml.Unmarshal([]byte(local), outStruct)
	}
	if group != DefaultGroupID {
		return Get(dataId, DefaultGroupID, outStruct)
	}
	return fmt.Errorf("no found config by:[dataId:%s, group:%s]", dataId, group)
}

func FromNacosService(dataId, group string) (string, error) {
	uri := fmt.Sprintf("%s/nacos/v1/cs/configs?dataId=%s&group=%s", instance.Host, dataId, group)
	from, err := helpers.HTTP().Get(uri)
	if err != nil {
		return "", err
	}
	return from, nil
}

func localCache(key string, result ...string) (string, error) {
	fileName := fmt.Sprintf("./runtime/%s.configCache", key)
	if len(result) > 0 { // 写入缓存数据
		if !helpers.FileIsExist("./runtime") && !helpers.CreateDir("./runtime") {
			return "", errors.New("create dir ./runtime Failed!")
		}
		return "", helpers.WriteToFile(fileName, result[0], false)
	}
	return helpers.FileRead(fileName)
}

func GetItem(dataKey, dataId, group string) (res interface{}, err error) {
	var outStruct map[string]interface{}
	err = Get(dataId, group, &outStruct)
	if err != nil {
		return nil, err
	}
	if dataKey != "." {
		keyLayer := strings.Split(dataKey, ".")
		deep := len(keyLayer) - 1
		for i, key := range keyLayer {
			if val, ok := outStruct[key]; ok {
				switch val.(type) {
				case map[interface{}]interface{}:
					tmpStr := make(map[string]interface{})
					for index, tmp := range val.(map[interface{}]interface{}) {
						tmpStr[index.(string)] = tmp
					}
					outStruct = tmpStr
					break
				case string:
					if i == deep {
						return val, nil
					}
					return nil, fmt.Errorf("beyond length limit: %s", key)
				case int:
					if i == deep {
						return val, nil
					}
					return nil, fmt.Errorf("beyond length limit: %s", key)
				case []interface{}:
					tmpStr := make(map[string]interface{})
					for index, tmp := range val.([]interface{}) {
						tmpStr[strconv.Itoa(index)] = tmp
					}
					outStruct = tmpStr
					break
				default:
					return nil, fmt.Errorf("no handle type: %s", fmt.Sprintf("%T", val))
				}
			} else {
				return nil, fmt.Errorf("not found key: %s, %+v", key, outStruct[key])
			}
		}
	}
	return outStruct, nil
}
