package middlewares

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"runtime"

	"mlj/pkg/common/consts"
	"mlj/pkg/response"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	if p := recover(); p != nil {
		errmsg := fmt.Sprintf("%s", p)
		pc, fn, line, _ := runtime.Caller(1)
		logbuf := &bytes.Buffer{}
		logbuf.WriteString(fmt.Sprintf("api(%s)[error] in %s[%s:%d] %v", c.Request.URL.Path, runtime.FuncForPC(pc).Name(), fn, line, errmsg))

		buf := make([]byte, 2048)
		buf = buf[:int(math.Min(float64(2048), float64(runtime.Stack(buf, true))))]
		logbuf.WriteString(fmt.Sprintf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf))

		log.Println(logbuf.String())
		if consts.EnvMode == consts.EnvModeProd {
			fmt.Println(logbuf.String())
		}

		response.Fail(c, errmsg)
	}
}
