
- 检测集群状态  
```
GET /_cat/health?v  
```

- 集群节点状态  
```
GET /_cat/nodes?v
```

- 索引状态  
```
GET /_cat/indices?v
```

- 创建索引  
```
PUT /index?pretty
{
    "settings": { ... any settings ... },
    "mappings": {
        "type_one": { ... any mappings ... },
        "type_two": { ... any mappings ... },
        ...
    }
}
```
- 修改索引
```
PUT /index/_settings
{
    "number_of_replicas": 1
}
```
- 删除索引
```
DELETE /index
DELETE /index_one,index_two
DELETE /index_*
DELETE /_all
```

```
# 查看所有索引
GET /_cat/indices?v
```

```
# 手动 refresh
POST /index/_refresh

#一般不需要手动执行，让 elasticsearch 自己来
PUT /index
{
  "settings": {
    "refresh_interval": "30s"
  }
}
```

- 新增记录  
```
POST /index/type/[id]?pretty
{
  "field_name":"field_value"
}

# 查询记录  
GET /customer/external/1?pretty
```

- 新增记录（不带 ID，随机分配）  
```
POST /index/type?pretty
{
  "field_name": "field_value"
}
```

- 更新数据（partial update）  
```
POST /index/type/[id]/_update?pretty
{
  "doc":{"field_name":"field_value"}
}
```

- 覆盖更新  
```
PUT /index/type/[id]?pretty
{
  "field_name":"field_value"
}
```

- 使用脚本更新  
```
# 内置脚本
POST /index/type/[id]/_update?pretty
{
  "script": "ctx._source.[field] += [value]"
}

# 外部脚本
# xxx 文件：ctx._source.tags+=tag_item
POST /index/type/[id]/_update?pretty
{
  "script": {
    "lang": "groovy", 
    "file": "xxx",
    "params": {
      "tag_item": "tag_value"
    }
  }
}
```

- 新增或更新（upsert）
> 如果指定的 document 不存在，就执行 upsert 中的初始化操作；  
> 如果指定的 document 存在，就执行 doc 或者 script 指定的 partial update 操作
```
POST /test_index/test_type/11/_update
{
   "script" : "ctx._source.num+=1",
   "upsert": {
       "num": 0,
       "tags": []
   }
}
```

- 删除文档  
```
DELETE /index/type/[id]?pretty

# 使用脚本删除
# xxx 文件：ctx.op = ctx._source.num == count ? 'delete' : 'none'
POST /index/type/[id]/_update?pretty
{
  "script": {
    "lang": "groovy",
    "file": "xxx",
    "params": {
      "count": 1
    }
  }
}
```

- 删除索引  
```
DELETE /index?pretty
```

- 批量处理(_bulk)
> bulk 操作中，任意一个操作失败，是不会影响其他的操作的，但是在返回结果里，会告诉你异常日志  
> 
> bulk request 会加载到内存里，如果太大的话，性能反而会下降，因此需要反复尝试一个最佳的 bulk size  
> 一般从 1000\~5000 条数据开始，尝试逐渐增加。另外，如果看大小的话，最好是在 5\~15MB 之间  
> 
> bulk api 对 json 的语法，有严格的要求，每个 json 串不能换行，只能放一行，同时一个 json 串和一个 json 串之间，必须有一个换行

```
# bulk 支持如下操作
#（1）delete：删除一个文档，只要 1 个 json 串就可以了
#（2）create：PUT /index/type/id/_create，强制创建
#（3）index：普通的 put 操作，可以是创建文档，也可以是全量替换文档
#（4）update：执行的 partial update 操作

POST /index/type/_bulk?pretty
{"index":{"_id":"[id]"}}
{"field_name":"field_value"}
{"create":{"_id":"[id]"}}
{"field_name":"field_value"}
{"update":{"_id":"[id]"}}
{"doc":{"field_name":"field_value"}}
{"delete":{"_id":"[id]"}}
# index 和 create 的区别：create 是文档不存在则创建，index 是创建新文档或替换已有文档
```

- 导入数据
```
# 与 _bulk 类似，创建 import.json
{"index":{"_id":"[id]"}}
{"field_name": "field_value"}

# 执行导入
curl -XPOST 'host:port/index/type/_bulk?pretty&refresh' --data-binary "@import.json"
```

- 创建索引
```
PUT /index
{
"settings": {...},
"mappings": {...}
}
```

- 更新索引配置
```
PUT /index/_settings
{
  ... [如 "number_of_replicas": 1]
}
```
