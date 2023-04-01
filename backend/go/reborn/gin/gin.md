
### Gin 框架简要使用
[Gin 中文文档](https://www.topgoer.cn/docs/ginkuangjia)  
[Gin 示例代码](https://github.com/gin-gonic/examples)

#### 获取参数
```go
// 路径中的
router.GET("/user/:name", func(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
})

// url 参数
firstname := c.DefaultQuery("firstname", "Guest")
lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") 的一种快捷方式

// 表单参数
name := c.PostForm("name")
message := c.PostForm("message")

// 获取 raw 参数
raw, err := ioutil.ReadAll(c.Request.Body)
```

#### 使用结构体来传递参数
```go
type User struct {
    Name  string `json:"name" form:"name" binding:"required"`
    Email string `json:"email" form:"email" binding:"required,email"`
}

func CreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBind(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 处理创建用户逻辑
    // ...

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
```

### 使用验证器来验证请求参数


#### 使用中间件来实现公共逻辑
跨域中间件
```go
import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

/**
 * 支持跨域
 */
func CrossDomain(c *gin.Context) {
	//origin
	allowOrigin := c.Request.Header.Get("Origin")
	if allowOrigin == "" {
		if referer := c.Request.Header.Get("Referer"); referer != "" {
			u, _ := url.ParseRequestURI(referer)
			allowOrigin = u.Host
		}
	}

	c.Header("Access-Control-Allow-Origin", allowOrigin)
	c.Header("Access-Control-Allow-Headers", "DNT, X-Mx-ReqToken, Keep-Alive, X-Requested-With, Cache-Control, If-Modified-Since, token,access-token, X-Origin, Origin, Accept, Content-Type, Referer, User-Agent, Cookie, access-token, crossdomain, withCredentials, authorization, XXX")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("x-container-pod", os.Getenv("HOSTNAME"))

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
```

登录校验
```go
// api登录授权
func AuthRequired(c *gin.Context) {
	u, err := GetUser(c)
	if err != nil {
		var redirect string
		if redirect = c.GetHeader("xxx-Redirect"); redirect == "" {
			if redirect = c.GetHeader("Referer"); redirect == "" {
				var scheme string
				if c.Request.TLS == nil {
					scheme = "http"
				} else {
					scheme = "https"
				}
				redirect = fmt.Sprintf("%s://%s%s", scheme, c.Request.Host, c.Request.RequestURI)
			}
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"result":   "autherror",
			"code":     "30010",
			"redirect": "https://xxx.com/?redirect=" + url.QueryEscape(redirect),
			"message":  "未登入授权: " + err.Error(),
			"data":     nil,
		})
		return
	}

	if u.Status != `on` {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"result":  "autherror",
			"code":    "30200",
			"message": "无权限操作",
			"data":    nil,
		})
		return
	}

	c.Set("user", u)
	c.Next()
}

func GetUser(c *gin.Context) (*User, error) {
	// xxx
}
```