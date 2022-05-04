/*
   块注释
*/
//单行注释

// @Title  文件名称
// @Description  文件描述
// @Author  yuanqi  ${DATE} ${TIME}
// @Update  yuanqi  ${DATE} ${TIME}
// package ${GO_PACKAGE_NAME}
package basic

// 单行引入
// import  "fmt"
// 多包引入，每包独占一行
// 使用绝对路径，避免使用相对路径

// User   用户对象，定义了用户的基础信息
type User struct {
	UserAge   int64  `json:"user_age"`    // 年龄
	Username  string `json:"user_name"`   // 用户名
	UserEmail string `json: "user_email"` // 邮箱
}

// @title 函数名称
// @description 函数描述
// @auth      yuanqi           ${DATE} ${TIME}
// @param     输入参数名        参数类型         "解释"
// @return    返回参数名        参数类型         "解释"
func (user User) test(userAge int64) {
	// 代码块的执行解释
	if userAge < 18 {

	}
}

func main() {

}
