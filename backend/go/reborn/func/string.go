package helper

import (
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// 转换为切片
func ConvertToSlice(str string, length int64) ([]string, error) {
	sliceRes := make([]string, length)
	if str != "" {
		if err := json.Unmarshal([]byte(str), &sliceRes); err != nil {
			return sliceRes, err
		}
	}
	return sliceRes, nil
}

// 转换为 map
func ConvertToMap(str string, length int64) ([]map[string]string, error) {
	mapRes := make([]map[string]string, length)
	if str != "" {
		if err := json.Unmarshal([]byte(str), &mapRes); err != nil {
			return mapRes, err
		}
	}
	return mapRes, nil
}

// 掩码
func Mask(str string) string {
	strLen := len(str)
	if strLen < 2 {
		return str
	}

	// 对 1-2 位进行掩码处理
	if strLen < 6 {
		return fmt.Sprintf("%s***%s", str[:1], str[strLen-1:])
	}

	// 对 3-后4 位进行掩码处理
	maskedDigits := fmt.Sprintf("%s****%s", str[:3], str[strLen-4:])

	return maskedDigits
}

// 移除 html 标签
func RemoveHTMLTags(htmlStr string) string {
	if htmlStr == "" {
		return ""
	}

	// 将 HTML 解析成一个 HTML 文档对象
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return htmlStr
	}

	// 遍历 HTML 文档节点，去除所有的标签
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			n.Parent.RemoveChild(n)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(doc)

	// 将处理后的 HTML 文档对象转换为字符串
	var buf strings.Builder
	html.Render(&buf, doc)

	return buf.String()
}