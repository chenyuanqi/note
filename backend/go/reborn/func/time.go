package helper

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// 解析时间字符串
func ParseTime(t string) *time.Time {
	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	return &ti
}

// 解析日期字符串
func ParseDate(t string) *time.Time {
	ti, _ := time.ParseInLocation("2006-01-02", t, time.Local)
	return &ti
}

// 几天前
func PastTime(p time.Time) string {
	ts := int64(time.Since(p).Seconds())

	var (
		min   int64 = 60
		hour  int64 = 60 * min
		day   int64 = 24 * hour
		month int64 = 30 * day
		year  int64 = 12 * month
	)

	//年 365d
	// ts /=
	y := ts / year
	ts = ts % year
	if y > 0 {
		return fmt.Sprintf("%d年前", y)
	}

	//月
	m := ts / month
	ts = ts % month
	if m > 0 {
		return fmt.Sprintf("%d月前", m)
	}

	//日
	d := ts / day
	ts = ts % day
	if d > 0 {
		return fmt.Sprintf("%d天前", d)
	}

	yms := (time.Duration(ts) * time.Second).String()
	yms = strings.NewReplacer("h", "时", "m", "分", "s", "秒").Replace(yms)
	return yms + "前"
}