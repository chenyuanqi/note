package helpers

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/**
 * 判定文件是否存在
 * true 存在 false 不存在
 */
func FileIsExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

/**
 * 创建文件夹
 * true 存在 false 不存在
 */
func CreateDir(dirName string) bool {
	return os.Mkdir(dirName, os.ModePerm) == nil
}

/**
  写入文件
  append 文件是否为追加写入，
  true:追加  false:覆盖
*/
func WriteToFile(fileName, content string, append bool) error {
	var flag int
	var f *os.File
	var err error
	if append {
		flag = os.O_WRONLY | os.O_CREATE
	} else {
		flag = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	}
	if FileIsExist(fileName) {
		f, err = os.OpenFile(fileName, flag, 0644)
	} else {
		f, err = os.Create(fileName)
	}
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.WriteString(f, content); err != nil {
		return err
	}
	return nil
}

// 读取文件
func FileRead(fileName string) (string, error) {
	if !FileIsExist(fileName) {
		return "", errors.New("file not exist")
	}
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return strings.Replace(string(contents), "\n", "", 1), nil
}

// 将时间字符串转换为时间戳
func ParseDataToTimestamp(timeStr string, flag int) int64 {
	var format string
	loc, _ := time.LoadLocation("Local")
	switch flag {
	case 1:
		format = "2006-01-02 15:04:05"
	case 2:
		format = "2006-01-02 15:04"
	case 3:
		format = "2006-01-02"
	case 4:
		format = "2006.01.02"
	default:
		format = "2006.01.02 15:04:05"
	}
	tt, _ := time.ParseInLocation(format, timeStr, loc)
	return tt.Unix()
}

func CkStr(str, defaultStr string) string {
	if str != "" {
		return str
	}
	return defaultStr
}

func Sum(content string) (s string) {
	if content == "" {
		content = time.Now().String()
	}
	s = fmt.Sprintf("%x", md5.Sum([]byte(content)))
	return
}

//19位 秒10位+微秒6位+随机数3位
func RedisScore(t time.Time) int64 {
	if t == (time.Time{}) {
		t = time.Now()
	}
	p12 := t.UnixNano()

	rand.Seed(t.UnixNano())
	p3 := rand.Int63n(100)

	return p12 + p3
}

func RedisScore2Time(score int64) time.Time {
	sec := score / 1e9
	nsec := score % 1e9
	return time.Unix(sec, nsec)
}
