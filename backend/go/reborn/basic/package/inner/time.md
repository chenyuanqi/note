
### time package：时间和日期
时间和日期是我们开发中经常会用到的，Go语言中的 time 包提供了时间显示和测量等所用的函数。  
时间一般包含时间值和时区，可以从 Go 语言中 time 包的源码中看出。
```go
type Time struct {
    // wall and ext encode the wall time seconds, wall time nanoseconds,
    // and optional monotonic clock reading in nanoseconds.
    //
    // From high to low bit position, wall encodes a 1-bit flag (hasMonotonic),
    // a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
    // The nanoseconds field is in the range [0, 999999999].
    // If the hasMonotonic bit is 0, then the 33-bit field must be zero
    // and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
    // If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
    // unsigned wall seconds since Jan 1 year 1885, and ext holds a
    // signed 64-bit monotonic clock reading, nanoseconds since process start.
    wall uint64 // 表示距离公元 1 年 1 月 1 日 00:00:00UTC 的秒数
    ext  int64 // 表示纳秒
    // loc specifies the Location that should be used to
    // determine the minute, hour, month, day, and year
    // that correspond to this Time.
    // The nil location means UTC.
    // All UTC times are represented with loc==nil, never loc==&utcLoc.
    loc *Location // 代表时区，主要处理偏移量，不同的时区，对应的时间不一样
}
```
Go 语言提供了 LoadLocation 方法和 FixedZone 方法来获取时区变量。
```go
FixedZone(name string, offset int) *Location
LoadLocation(name string) (*Location, error) // name 为时区的名字
```

**时间的获取**  
1) 获取当前时间  
通过 time.Now() 函数来获取当前的时间对象，然后通过事件对象来获取当前的时间信息。  
```go
package main
import (
    "fmt"
    "time"
)
func main() {
    now := time.Now() //获取当前时间
    fmt.Printf("current time:%v\n", now) // current time:2019-12-12 12:33:19.4712277 +0800 CST m=+0.006980401
    year := now.Year()     //年
    month := now.Month()   //月
    day := now.Day()       //日
    hour := now.Hour()     //小时
    minute := now.Minute() //分钟
    second := now.Second() //秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) // 2022-12-12 12:33:19
}
```
2) 获取时间戳  
时间戳是自 1970 年 1 月 1 日（08:00:00GMT）至当前时间的总毫秒数，它也被称为 Unix 时间戳（UnixTimestamp）。
```go
package main
import (
    "fmt"
    "time"
)
func main() {
    // 基于时间对象获取时间戳
    now := time.Now()            //获取当前时间
    timestamp1 := now.Unix()     //时间戳
    timestamp2 := now.UnixNano() //纳秒时间戳
    fmt.Printf("现在的时间戳：%v\n", timestamp1)
    fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)

    // 使用 time.Unix() 函数可以将时间戳转为时间格式
    timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
    fmt.Println(timeObj)
    year := timeObj.Year()     //年
    month := timeObj.Month()   //月
    day := timeObj.Day()       //日
    hour := timeObj.Hour()     //小时
    minute := timeObj.Minute() //分钟
    second := timeObj.Second() //秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
```
3) 获取当前是星期几  
time 包中的 Weekday 函数能够返回某个时间点所对应是一周中的周几。
```go
package main
import (
    "fmt"
    "time"
)
func main() {
    //时间戳
    t := time.Now()
    fmt.Println(t.Weekday().String()) // Thursday
}
```

**时间操作函数**  
1) Add  
某个时间 + 时间间隔，Add 函数可以返回时间点 t + 时间间隔 d 的值。
```go
func (t Time) Add(d Duration) Time

package main
import (
    "fmt"
    "time"
)
func main() {
    now := time.Now()
    later := now.Add(time.Hour) // 当前时间加1小时后的时间
    fmt.Println(later)
}
```
2) Sub  
求两个时间之间的差值。返回一个时间段 t - u 的值。如果结果超出了 Duration 可以表示的最大值或最小值，将返回最大值或最小值，要获取时间点 t - d（d 为 Duration），可以使用 t.Add(-d)。  
```go
func (t Time) Sub(u Time) Duration
```
3) Equal  
判断两个时间是否相同，Equal 函数会考虑时区的影响，因此不同时区标准的时间也可以正确比较，Equal 方法和用 t==u 不同，Equal 方法还会比较地点和时区信息。
```go
func (t Time) Equal(u Time) bool
```
4) Before  
判断一个时间点是否在另一个时间点之前，如果 t 代表的时间点在 u 之前，则返回真，否则返回假。
```go
func (t Time) Before(u Time) bool
```
5) After  
判断一个时间点是否在另一个时间点之后，如果 t 代表的时间点在 u 之后，则返回真，否则返回假。
```go
func (t Time) After(u Time) bool
```

**定时器**  
使用 time.Tick(时间间隔) 可以设置定时器，定时器的本质上是一个通道（channel）。  
```go
package main
import (
    "fmt"
    "time"
)
func main() {
    ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
    for i := range ticker {
        fmt.Println(i) //每秒都会执行的任务
    }
}
```

**时间格式化**  
时间类型有一个自带的 Format 方法进行格式化，需要注意的是 Go 语言中格式化时间模板不是常见的 Y-m-d H:M:S 而是使用 Go 语言的诞生时间 2006 年 1 月 2 号 15 点 04 分 05 秒。  
`提示：如果想将时间格式化为 12 小时格式，需指定 PM。`
```go
package main
import (
    "fmt"
    "time"
)
func main() {
    now := time.Now()
    // 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
    // 24小时制
    fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
    // 12小时制
    fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
    fmt.Println(now.Format("2006/01/02 15:04"))
    fmt.Println(now.Format("15:04 2006/01/02"))
    fmt.Println(now.Format("2006/01/02"))
}
```

**解析字符串格式的时间**   
ParseInLocation 与 Parse 函数类似，但有两个重要的不同之处：  
- 第一，当缺少时区信息时，Parse 将时间解释为 UTC 时间，而 ParseInLocation 将返回值的 Location 设置为 loc；
- 第二，当时间字符串提供了时区偏移量信息时，Parse 会尝试去匹配本地时区，而 ParseInLocation 会去匹配 loc。
```go
// Parse 函数可以解析一个格式化的时间字符串并返回它代表的时间
func Parse(layout, value string) (Time, error)

// ParseInLocation 函数
func ParseInLocation(layout, value string, loc *Location) (Time, error)

package main
import (
    "fmt"
    "time"
)
func main() {
    var layout string = "2006-01-02 15:04:05"
    var timeStr string = "2019-12-12 15:22:12"
    timeObj1, _ := time.Parse(layout, timeStr)
    fmt.Println(timeObj1)
    timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
    fmt.Println(timeObj2)
}
```
