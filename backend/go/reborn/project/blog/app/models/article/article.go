package article

import (
	"blog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	ID      uint64
	Title   string
	Content string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}
