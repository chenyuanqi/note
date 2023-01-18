package helpers

import (
	"encoding/json"
	"fmt"
)

type dbSave struct {
	Output output   `json:"output"`
	Data   saveData `json:"data"`
}
type output struct {
	Name  string `json:"name"`  // 数据库名
	Table string `json:"table"` // 数据表名
}
type saveData struct {
	Insert map[string]interface{} `json:"insert"` // 插入数据
	Update map[string]interface{} `json:"update"` // 更新数据
	Filter map[string]interface{} `json:"filter"` // 过滤数据
}

type DBSave struct {
	ConnName   string
	TableName  string
	InsertData map[string]interface{}
	UpdateData map[string]interface{}
	FilterData map[string]interface{}
}

var dbSaveKey = "uni-dbsave-list"

func (d DBSave) Save() bool {
	data := dbSave{
		Output: output{
			Name:  d.ConnName,
			Table: d.TableName,
		},
		Data: saveData{
			Insert: d.InsertData,
			Update: d.UpdateData,
			Filter: d.FilterData,
		},
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("DBSave has error: %+v", err)
		//Log().Error("dbsave err ", err)
		return false
	}
	success, err := RedisClient(CCRedisConf("config")).RPush(dbSaveKey, string(dataJson)).Result()
	if err != nil {
		fmt.Printf("DBSave has error: %+v", err)
		return false
	}
	return success == 1
}
