
### time 包
time 包为我们提供了一个数据类型 time.Time（作为值使用）以及显示和测量时间和日期的功能函数。  
```golang
type Time struct {
    wall uint64 // 表示距离公元 1 年 1 月 1 日 00:00:00UTC 的秒数
    ext  int64 // 表示纳秒
    loc *Location // 代表时区，主要处理偏移量，不同的时区，对应的时间不一样
}
```

当前时间可以使用 time.Now() 获取，或者使用 t.Day()、t.Minute() 等等来获取时间的一部分；你甚至可以自定义时间格式化字符串，例如： fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) 将会输出 21.07.2011。  

Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。Location 类型映射某个时区的时间，UTC 表示通用协调世界时间。  

包中的一个预定义函数 func (t Time) Format(layout string) string 可以根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串，你可以使用一些预定义的格式，如：time.ANSIC 或 time.RFC822。一般的格式化设计是通过对于一个标准时间的格式化描述来展现的。  
```golang
fmt.Println(t.Format("02 Jan 2006 15:04"))  // 21 Jul 2011 10:31
```

如果你需要在应用程序在经过一定时间或周期执行某项任务（事件处理的特例），则可以使用 time.After 或者 time.Ticker。  

time.Sleep（d Duration） 可以实现对某个进程（实质上是 goroutine）时长为 d 的暂停。  


### time 包的使用

```golang
t := time.Now()
fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) // 21.12.2011

t = time.Now().UTC()
fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011

now := time.Now() //获取当前时间
fmt.Printf("current time:%v\n", now)
year := now.Year()     //年
month := now.Month()   //月
day := now.Day()       //日
hour := now.Hour()     //小时
minute := now.Minute() //分钟
second := now.Second() //秒
fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
fmt.Println(now.Weekday().String()) // 周几

now := time.Now()            //获取当前时间
timestamp1 := now.Unix()     //时间戳
timestamp2 := now.UnixNano() //纳秒时间戳
fmt.Printf("现在的时间戳：%v\n", timestamp1)
fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)

// calculating times:
week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
week_from_now := t.Add(time.Duration(week))
fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011

// formatting times:
fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011

// The time must be 2006-01-02 15:04:05
fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
s := t.Format("20060102")
fmt.Println(t, "=>", s) // Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
```
