
- 查询全部  
```
GET /index/_search
{
  "query": {"match_all": {}}
}
```

- 分页查询  
```
GET /index/_search
{
  "query": {"match_all": {}}, 
  "from": [number],
  "size": [number]
}
```

- 排序  
```
GET /index/_search
{
  "query": {"match_all": {}},
  "sort": [
    {"field_name":"asc/desc"}
  ]
}
```

- 指定字段  
```
GET /index/_search
{
  "query": {"match_all": {}},
  "_source": ["field_name", "field_name"]
}
```

- highlight search（高亮搜索结果）
```
GET /index/type/_search
{
    "query" : {
        "match" : {
            "field_name": "field_value"
        }
    },
    "highlight": {
        "fields" : {
            "field_name" : {}
        }
    }
}
```

- 匹配查询
```
# field_value 为数值型则精确匹配，字符型则包含匹配（**空格代表逻辑或**）
GET /index/_search
{
  "query": {"match": {
    "field_name": "field_value"
  }}
}
```

- 条件组合
```
# bool 查询，should 相当于逻辑或，must 相当于逻辑与，must_not 相当于逻辑非
GET /index/_search
{
  "query": {
    "bool": {
      "should": [
        {"match": {
          "field_name": "field_value"
        }},
        {"match": {
          "field_name": "field_value"
        }}
      ]
    }
  }
}
```

- 区间查询
```
# include_lower 是否包含小于等于，include_upper 是否包含大于等于
GET /index/_search
{
  "query": {
    "range": {
      "field_name": {
        "gte": "field_value",
        "lte": "field_value",
        "include_lower": false/true,
        "include_upper": false/true
      }
    }
  }
}
```

- 前缀匹配
```
GET /index/type/_search
{
  "query": {
    "prefix": {
      "field_name": {
        "value": "field_value"
      }
    }
  }
}
```

- 模糊查询
```
# 通配符 * 代表一个或多个字符，? 仅代表一个
GET /index/type/_search
{
  "query": {
    "wildcard": {
      "field_name": {
        "value": "field_value"
      }
    }
  }
}
```

```
# boost 权重配置，min_similarity 最小相似度，prefix_length 分词项共同前缀长度
GET /index/type/_search
{
  "query": {
    "fuzzy": {
      "field_name": {
        "value": "field_value",
        "boost": [default: 1.0],
        "min_similarity": [default: 0.5, value: 0~1 / 2d etc.],
        "prefix_length": [default: 0]
      }
    }
  }
}
```

- mget 查询
```
GET /index/type/_search
{
  "docs": [
    { "_id": [number] }
  ]
}

# 多个 id
GET /index/type/_search
{
  "ids": [numbers, numbers]
}
```

- term 查询
```
# 处理匹配查询，也可查询某数值在数组中
GET /index/type/_search
{
   "query": {
      "term": {
         "field_name": "field_value"
      }
   }
}
```

- 单条数据
```
# pretty 即美化格式输出
GET /index/type/[id]?pretty
```

- query string
```
# query_syntax 语法说明
# ① query_syntax 值为一般字符串，请求执行全文检索
# ② query_syntax 值为 field_name:string 根据指定字段全文检索（全匹配）
# ③ query_syntax 值为 field_name:'string' 指定字段精确检索
# ④ query_syntax 值如 name:"production" AND date:"2017/08/08" ，执行多个检索条件组合，使用 NOT、AND、OR，注意必须是大写的，并且条件与条件之间需要空格分隔
# ⑤ query_syntax 值如 name:product* ，使用通配符? 表示单字母或* 表示任意个字母
# ⑥ query_syntax 值如 order:product[1-9] ，使用正则，性能较差（[参考链接](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-regexp-query.html#regexp-syntax)）
# ⑦ query_syntax 值如 name:foo~ ，用 ~ 表示搜索单词可能有一个或两个字母写得不对，按照相似度返回结果
# ⑧ query_syntax 值如 created_at:["now-12h" TO "now"] ，对于数值（如 age:>18 ）和时间进行范围搜索
GET /index/type/_search?q=[query_syntax]

# 或者这样
GET /index/type/_search
{
   "query": {
      "query_string": {
         "query": "[query_syntax]"
      }
   }
}
```

- full-text search（全文检索）
> 将输入的搜索串拆解开来，去倒排索引里面去一一匹配，只要能匹配上任意一个拆解后的单词，就可以作为结果返回  
```
GET /index/type/_search
{
    "query" : {
        "match" : {
            "field_name" : "field_value"
        }
    }
}
```

- phrase search（短语搜索）
> 要求输入的搜索串，必须在指定的字段文本中，完全包含一模一样的，才可以算匹配，才能作为结果返回  
```
GET /index/type/_search
{
    "query" : {
        "match_phrase" : {
            "field_name" : "field_value"
        }
    }
}
```

- match phrase prefix（短语前缀匹配）
> 短语搜索，有遗漏时考虑使用  
```
POST /index/type/_search
{
  "profile": "true",
  "query": {
    "match_phrase_prefix": {
      "field_name": {
        "query": "field_value",
        "max_expansions": [expansion_number]
      }
    }
  }
}
```

- 包含查询
```
# 类似于 whereIn，多项则重复 match 部分
# mode（must => &&, should => ||）
GET /index/type/_search
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "bool": {
                "should": [{
                  "match": {
                    [field_name]: {
                      "query": [field_value],
                      "type": "phrase"
                    }
                  }
                }]
              }
            }
          ]
        }
      }
    }
  }
}
```

- 不包含查询
```
# 类似于 whereNotIn，多项则重复 match 部分
# mode（must => &&, should => ||）
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "bool": {
                "must_not": {
                  "bool": {
                    "should": [
                      {
                        "match": {
                          [field_name]: {
                            "query": [field_value],
                            "type": "phrase"
                          }
                        }
                      }
                    ]
                  }
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

- 条件查询
```
# 类似于 where field_name = field_value
# mode（must => &&, should => ||）
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "match": {
                [field_name]: {
                  "query": [field_value],
                  "type": "phrase"
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

```
# 类似于 where field_name != field_value 或 where field_name <> field_value
# mode（must => &&, should => ||）
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "bool": {
                "must_not": {
                  "match": {
                    [field_name]: {
                      "query": [field_value],
                      "type": "phrase"
                    }
                  }
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

```
# 类似于 where field_name > field_value（如果是 >= ，include_lower 为 true）
# mode（must => &&, should => ||）
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "range": {
                [field_name]: {
                  "from": [field_value],
                  "to": null,
                  "include_lower": [true/false],
                  "include_upper": true
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

```
# 类似于 where field_name < field_value（如果是 <= ，include_upper 为 true）
# mode（must => &&, should => ||）
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "range": {
                [field_name]: {
                  "from": null,
                  "to": [field_value],
                  "include_lower": true,
                  "include_upper": [true/false]
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

```
# 类似于 where field_name like field_value
# field_value 中 _ 代表 0 个或 1 个，% 代表 0 个或多个
# mode（must => &&, should => ||）
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "wildcard": {
                [field_name]: [field_value]
              }
            }
          ]
        }
      }
    }
  }
}
```

- 区间条件查询
```
# 类似于 whereBetween
# mode（must => &&, should => ||）
# include_lower 是否包含小于等于，include_upper 是否包含大于等于
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "range": {
                [field_name]: {
                  "from": [field_value1],
                  "to": [field_value2],
                  "include_lower": [true/false],
                  "include_upper": [true/false],
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

- 不在区间条件查询
```
# 类似于 whereNotBetween
# mode（must => &&, should => ||）
# include_lower 是否包含小于等于，include_upper 是否包含大于等于
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "bool": {
                "must_not": {
                  "range": {
                    [field_name]: {
                      "from": [field_value1],
                      "to": [field_value2],
                      "include_lower": [true/false],
                      "include_upper": [true/false],
                    }
                  }
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

- 字段为 NULL 查询
```
# 类似于 isNull
# mode（must => &&, should => ||）
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "missing": {
                "field": [field_name]
              }
            }
          ]
        }
      }
    }
  }
}
```

- 字段不为 NULL 查询
```
# 类似于 isNotNull
# mode（must => &&, should => ||）
{
  "query": {
    "bool": {
      "must": {
        "bool": {
          [mode]: [
            {
              "bool": {
                "must_not": {
                  "missing": {
                    "field": [field_name]
                  }
                }
              }
            }
          ]
        }
      }
    }
  }
}
```

- 过滤查询
```
{
  "filter": {
    "term": {
      "field_name": "field_value"
    }
  }
}
```

- 滚动查询
> 滚动的 ID，scroll 时间过期或重启 elastic 后失效  
```
# search_type 为 scan 时代表扫描，scroll 时代表滚动
# scroll 代表保持查询的过期时间
# query 部分为任意查询条件
# size 为查询记录数
GET /index/type/_search?search_type=scan&scroll=1m
{
  "query":{
    "match_all":{}
  },
  "size": [size_number]
}
```

```
# 根据 scroll_id 查询
GET /index/type/_search/scroll
{
  "scroll": 1m,
  "scroll_id": [scroll_id]
}
```