package article

import (
	"blog/app/models"
	"blog/app/models/user"
	"blog/pkg/route"

	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title   string `gorm:"type:varchar(255);not null;" valid:"title"`
	Content string `gorm:"type:longtext;not null;" valid:"content"`
	UserID  uint64 `gorm:"not null;index"`
	User    user.User
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}
