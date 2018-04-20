
- 统计
```
# 统计查询中某字段出现次数
# aggs 也可以是 aggregations
GET /index/type/_search
{
  "aggs": {
    "count_field_name": {
      "terms": { "field": "field_name" }
    }
  }
}
```

- 最大值
```
# 统计查询中某字段最大值
# aggs 也可以是 aggregations
GET /index/type/_search
{
  "aggs": {
    "max_field_name": {
      "max": { "field": "field_name" }
    }
  }
}
```

- 最小值
```
# 统计查询中某字段最小值
# aggs 也可以是 aggregations
GET /index/type/_search
{
  "aggs": {
    "min_field_name": {
      "min": { "field": "field_name" }
    }
  }
}
```

- 求和
```
# 统计查询中某字段求和
# aggs 也可以是 aggregations
GET /index/type/_search
{
  "aggs": {
    "sum_field_name": {
      "sum": { "field": "field_name" }
    }
  }
}
```

- 求平均值
```
# 统计查询中某字段的平均值
GET /index/type/_search
{
  "aggs": {
    "avg_field_name": {
      "avg": { "field": "field_name" }
    }
  }
}
```
