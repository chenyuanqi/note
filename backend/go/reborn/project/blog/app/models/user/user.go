package user

import (
	"blog/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	// Name     string `gorm:"column:name;type:varchar(255);not null;unique"`
	// Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique;"`
	// Password string `gorm:"column:password;type:varchar(255)"`

	// GORM 默认会将键的小写化作为字段名称，column 项可去除，并且默认是允许 NULL
	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段
	PasswordConfirm string ` gorm:"-" valid:"password_confirm"`
}

// Link 方法用来生成用户链接
func (user User) Link() string {
	return ""
}
