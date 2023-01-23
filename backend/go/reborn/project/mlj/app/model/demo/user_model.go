//Package category 模型
package demo

import (
	"crypto/md5"
	"fmt"
	"time"

	"mlj/pkg/database"

	"gorm.io/gorm"
)

type Users struct {
	ID        int64     `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"-"`
}

func (u *Users) Find() []Users {
	result := []Users{}
	database.DB.Find(&result)

	return result
}

func (u *Users) FindByID(id int64) {
	database.DB.First(&u, id)
}

func (u *Users) AfterFind(tx *gorm.DB) (err error) {
	return
}

func (u *Users) BeforeCreate(scope *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	data := []byte(u.Password)
	has := md5.Sum(data)
	u.Password = fmt.Sprintf("%x", has)

	return
}

func (u *Users) Create() {
	database.DB.Create(&u)
}

func (u *Users) Save() (rowsAffected int64) {
	result := database.DB.Save(&u)
	return result.RowsAffected
}

func (u *Users) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&u)
	return result.RowsAffected
}

func (u Users) TableName() string {
	return "users"
}

func (u Users) CreatedAtDateTime() string {
	return u.CreatedAt.Format("2006-01-02 15:04:05")
}

func (u Users) UpdatedAtDateTime() string {
	return u.UpdatedAt.Format("2006-01-02 15:04:05")
}
