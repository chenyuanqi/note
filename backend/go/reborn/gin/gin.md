
### Gin 框架简要使用
[Gin 中文文档](https://www.topgoer.cn/docs/ginkuangjia)  
[Gin 示例代码](https://github.com/gin-gonic/examples)

#### 重定向
```go
ginServer.GET("/test", func(ctx *gin.Context) {
	// 重定向 -> 301
	ctx.Redirect(301, "https://xxx/")
})
```

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

#### 使用验证器来验证请求参数
在 gin 框架中，可以使用验证器来验证请求参数。常用的验证器有：  
1. binding：将请求参数绑定到结构体上，并进行验证  
2. validator：对结构体中的字段进行验证  
```go
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required,email"`
	Age   int    `form:"age" binding:"gte=18,lte=100"`
}

func main() {
	r := gin.Default()

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkName", checkName)
	}

	r.GET("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.Run(":8080")
}

// 自定义验证函数，用于验证 Name 字段
func checkName(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	return name != "admin"
}
```
在上述代码中，我们定义了一个 User 结构体，并使用 binding 和 validator 对其进行验证。在路由处理函数中，我们使用 ShouldBind 方法将请求参数绑定到 User 结构体上，并进行验证。如果验证失败，我们返回一个包含错误信息的 JSON 响应；如果验证成功，我们返回一个包含 User 结构体的 JSON 响应。  
在注册验证器时，我们还自定义了一个 checkName 函数，用于验证 Name 字段的值是否为 "admin"。这个函数会在验证 Name 字段时被调用。  
需要注意的是，我们在注册验证器时使用了 binding.Validator.Engine() 方法，该方法返回的是一个 validator 的实例，我们可以通过类型断言将其转换为 *validator.Validate 类型，从而使用其 RegisterValidation 方法注册自定义验证函数。

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