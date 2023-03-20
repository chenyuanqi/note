package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func Query() error {
	type PidItem struct {
		Uid       string
		Platform  string
		Account   string
		Pid       string
		WeiboName string
	}
	var (
		pids     []*PidItem
		querysql = fmt.Sprintf(`select uid, platform, account, pid, weibo_name from ho_pids where status = '%s' and default_flag = %d`, "available", 1)
	)

	db, _ := GetDB()
	if err := db.Raw(querysql).Scan(&pids).Error; err != nil {
		return fmt.Errorf("query err: %s", err)
	}

	return nil
}

type Xxx struct {
	Id        uint64
	Date      time.Time
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func BatchInsert(data map[string]*Xxx) {
	var insertFields = []string{
		"date",
		"category",
		"created_at",
		"updated_at",
	}

	insertSqlTemplate := fmt.Sprintf("insert into xxx(`%s`) values", strings.Join(insertFields, "`, `"))
	batch := make([]string, 0)
	for _, item := range data {
		row := []string{
			item.Date.Format("2006-01-02"),
			item.Category,
			item.CreatedAt.Format("2006-01-02 15:04:05"),
			item.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		temp := fmt.Sprintf("('%s')", strings.Join(row, "', '"))
		batch = append(batch, temp)
	}
	log.Printf("len = %d", len(batch))
	if len(batch) == 0 {
		log.Fatal("batch empty")
	}

	sql := fmt.Sprintf(insertSqlTemplate, strings.Join(batch, `, `))
	log.Printf("sql: %s", sql)

	db, _ := GetDB()
	res := db.Exec(sql)
	if res.Error != nil {
		log.Fatalf("sql err: %s", res.Error)
	}
	log.Printf("affected %d", res.RowsAffected)
}

func Update(ids []int) {
	now := time.Now()
	upts := map[string]interface{}{
		"category":   "xx",
		"updated_at": now.Format("2006-01-02 15:04:05"),
	}

	db, _ := GetDB()
	res := db.Model(&Xxx{}).Where("id in ?", ids).UpdateColumns(upts)
	if err := res.Error; err != nil {
		log.Fatalf("ids update err: %v", err)
	}
	log.Printf("db affected %d", res.RowsAffected)
}

func Upsert(data []map[string]interface{}) {
	db, _ := GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				// 重复的字段
				{Name: "date"},
				{Name: "category"},
				{Name: "created_at"},
				{Name: "updated_at"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"category",
				"updated_at",
			}),
		}).CreateInBatches(&data, len(data)).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalf("items db trans err: %v", err)
	}
}

func Delete(date string) error {
	sql := fmt.Sprintf(`delete from xxx where date = "%s"`, date)
	db, _ := GetDB()
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	return nil
}

func Transaction() {
	db, _ := GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		Delete("2023-01-01")

		return nil
	})
	if err != nil {
		log.Fatalf("db transaction err: %s", err)
	}
}

// 实现 MarshalJSON 方法
func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User // 定义别名

	return json.Marshal(&struct {
		CreatedAt time.Time `json:"created_at"`
		*Alias    // 使用别名
	}{
		CreatedAt: u.CreatedAt.Local(),
		Alias:     (*Alias)(u),
	})
}

func main() {}
