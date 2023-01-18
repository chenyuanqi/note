package helpers

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"mlj/pkg/common/consts"

	"gopkg.in/yaml.v2"
)

var configServiceHost string

const ConfigServiceDev = "http://config.maliujia"
const ConfigService = "http://nacos-k8s.default"

func init() {
	configServiceHost = ConfigServiceDev
	if consts.EnvMode == consts.EnvModeProd {
		configServiceHost = ConfigService
	}
}

type ConfigCenter struct {
	Group string
}

var CCDefault = &ConfigCenter{}

func (c *ConfigCenter) Get(dataId, group string, outStruct interface{}) error {
	if group == "" {
		if c.Group != "" {
			group = c.Group
		} else {
			group = "DEFAULT_GROUP"
		}
	}
	key := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s", dataId, group))))
	//尝试内存缓存中取值
	result, err := Cache().Get(key, 0)
	if err != nil || result == "" {
		// Log().Info("Cache not found key/value")
		fmt.Println("Cache not found key/value")
	} else {
		return json.Unmarshal([]byte(result), outStruct)
	}

	//尝试CC取值
	result, err = c.FromNacosService(dataId, group)
	if err != nil {
		// Log().Error("config center err:", err)
		fmt.Println("config center err:", err)
	} else {
		if err := yaml.Unmarshal([]byte(result), outStruct); err != nil {
			return err
		}
		jsonRes, _ := json.Marshal(outStruct)
		strRes := string(jsonRes)
		Cache().Set(key, strRes, 60*time.Second) //缓存到内存表
		//缓存到本地
		if _, err = c.LocalCache(key, strRes); err != nil {
			// Log().Error(err)
			fmt.Println(err)
		}
		return nil
	}

	//尝试本地文件读取
	result, err = c.LocalCache(key)
	if err != nil {
		//写日志
		// Log().Info("try read localCache err: ", err)
		fmt.Println("try read localCache err: ", err)
	} else {
		return json.Unmarshal([]byte(result), outStruct)
	}
	if group != "DEFAULT_GROUP" {
		return c.Get(dataId, "DEFAULT_GROUP", outStruct)
	}
	err = errors.New(fmt.Sprintf("no found config by:[dataId:%s, group:%s]", dataId, group))
	// Log().Error(err.Error())
	fmt.Println(err.Error())
	return err
}

func (c *ConfigCenter) FromNacosService(dataId, group string) (string, error) {
	from, err := HTTP().Get(fmt.Sprintf("%s/nacos/v1/cs/configs?dataId=%s&group=%s", configServiceHost, dataId, group))
	if err != nil {
		return "", err
	}
	return from, nil
}

func (c *ConfigCenter) LocalCache(key string, result ...string) (string, error) {
	fileName := fmt.Sprintf("./runtime/%s.configCache", key)
	if len(result) > 0 { // 写入缓存数据
		if !FileIsExist("./runtime") && !CreateDir("./runtime") {
			return "", errors.New("Create Dir [runtime] Failed!")
		}
		return "", WriteToFile(fileName, result[0], false)
	}
	return FileRead(fileName)
}
