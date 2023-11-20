
### JsonTime 的封装
```go
package model

import (
	"database/sql/driver"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

type JsonTime time.Time

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return
}

// MarshalJSON echo api json response
func (t JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

// for sql log, print readable format
func (t JsonTime) String() string {
	if t == JsonTime(time.Time{}) {
		return ""
	}

	ts := time.Time(t)
	return ts.Format(TimeFormat)
}

// 转成 time.Time
func (t JsonTime) ToTime() time.Time {
	return time.Time(t)
}

// 判断是否为空
func (t JsonTime) IsEmpty() bool {
	return t == JsonTime(time.Time{})
}

// 指定格式输出
func (t JsonTime) ToCustomDateString(format string) string {
	if t == JsonTime(time.Time{}) {
		return ""
	}

	if format == "" {
		format = "2006-01-02"
	}

	ts := time.Time(t)
	return ts.Format(format)
}
func (t JsonTime) ToCustomDateNoYearString(format string) string {
	if t == JsonTime(time.Time{}) {
		return ""
	}

	if format == "" {
		format = "01-02"
	}

	ts := time.Time(t)
	return ts.Format(format)
}

func (t JsonTime) ToDateString() string {
	if t == JsonTime(time.Time{}) {
		return ""
	}

	ts := time.Time(t)
	return ts.Format("2006-01-02")
}

func (t JsonTime) ToDateTimeString() string {
	if t == JsonTime(time.Time{}) {
		return ""
	}

	ts := time.Time(t)
	return ts.Format("2006-01-02 15:04:05")
}

// insert into database conversion
func (t JsonTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// read from database conversion
func (t *JsonTime) Scan(src interface{}) error {
	if val, ok := src.(time.Time); ok {
		*t = JsonTime(val)
	}
	return nil
}
```

### JsonTime 的使用
1、和 time.Time 的转换
```go
*model.JsonTime字段 = (*model.JsonTime)(&time.Time字段)
t := model.JsonTime(time.Time/model.JsonTime字段)
*model.JsonTime字段 = &t

// 避免空指针
var xxDate *JsonTime
if item.xxDate != nil {
	convertedTime := JsonTime(*item.xxDate)
	xxDate = &convertedTime
}
```

2、赋值当前时间
```go
model.JsonTime字段 = model.JsonTime(time.Now())
// 操作时间
model.JsonTime(time.Now().Add(10 * time.Minute))
```

3、字符串转换
```go
t, _ := time.ParseInLocation("2006-01-02", "xxxx-xx-xx", time.Local)
mt := model.JsonTime(t)
*model.JsonTime字段 = &mt
```

