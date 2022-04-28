package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type demo struct {
	Id       int
	Title    string `json:"title"`
	Content  string `json:"content"`
	DateTime string `json:"date_time"`
}

var db *sql.DB

func main() {
	db = connectToDb()
	createTable()
	launchServer()
}

// 连接到数据库
func connectToDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")
	if err != nil {
		panic(err)
	}
	// 设置可重用链接的最长时间（0为不限制）
	db.SetConnMaxLifetime(time.Hour * 1)
	// 设置连接到数据库的最大数量（默认值为0，即不限制）
	db.SetMaxOpenConns(5)
	// 设置空闲连接的最大数量（默认值为2）
	db.SetMaxIdleConns(5)
	fmt.Println("connect success")
	return db
}

// 创建数据表
func createTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `demo` (" +
		"`id` bigint(20) NOT NULL AUTO_INCREMENT," +
		"`title` varchar(45) DEFAULT ''," +
		"`content` varchar(45) DEFAULT ''," +
		"`date_time` varchar(45) DEFAULT ''," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;")
	if err != nil {
		panic(err)
	}
	fmt.Println("table exist or create success")
}

// 添加数据到数据库
func add(data demo) {
	stmtInsert, _ := db.Prepare("INSERT INTO demo SET title=?,content=?,date_time=?")
	res, _ := stmtInsert.Exec(data.Title, data.Content, data.DateTime)
	id, _ := res.LastInsertId()
	fmt.Printf("add success: %d\n", id)
}

// 删除一条数据
func del(id string) {
	stmtDelete, _ := db.Prepare("DELETE FROM demo WHERE id=?")
	res, _ := stmtDelete.Exec(id)
	rawsCount, _ := res.RowsAffected()
	fmt.Printf("delete success: %d\n", rawsCount)
}

// 更新数据到数据库
func update(id string, data demo) {
	stmtInsert, _ := db.Prepare("UPDATE demo SET title=?, content=?, date_time=? WHERE id=?")
	res, _ := stmtInsert.Exec(data.Title, data.Content, data.DateTime, id)
	rawsCount, _ := res.RowsAffected()
	fmt.Printf("update success: %d\n", rawsCount)
}

// 从数据库获取数据
func query(id string) []demo {
	var demos []demo
	var rows *sql.Rows
	if id == "" {
		rows, _ = db.Query("SELECT * FROM demo")
	} else {
		rows, _ = db.Query("SELECT * FROM demo WHERE id=" + id)
	}
	for rows.Next() {
		var singleNote demo
		rows.Scan(&singleNote.Id, &singleNote.Title, &singleNote.Content, &singleNote.DateTime)
		demos = append(demos, singleNote)
	}
	return demos
}

// 启动服务器
func launchServer() {
	//响应/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome\n")
		fmt.Fprintf(w, "▶▶ /add 添加新的数据\n")
		fmt.Fprintf(w, "▶▶ /delete 根据ID删除数据\n")
		fmt.Fprintf(w, "▶▶ /update 根据ID更新数据\n")
		fmt.Fprintf(w, "▶▶ /query 获取全部数据或根据ID获取单条数据\n")
	})

	//响应/add，从传入的参数新增一条记事本
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintln(w, "failed request")
			} else {
				title := r.FormValue("title")
				content := r.FormValue("content")
				dateTime := r.FormValue("date_time")
				add(demo{Title: title, Content: content, DateTime: dateTime})
				fmt.Fprintln(w, "add success")
			}
		}
	})

	//响应/delete，从传入的参数删除一条记事本
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "failed request")
		} else {
			id := r.FormValue("id")
			del(id)
			fmt.Fprintln(w, "delete success")
		}
	})

	//响应/update，更新一条数据
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "failed request")
		} else {
			id := r.FormValue("id")
			title := r.FormValue("title")
			content := r.FormValue("content")
			dateTime := r.FormValue("date_time")
			update(id, demo{Title: title, Content: content, DateTime: dateTime})
			fmt.Fprintln(w, "update success")
		}
	})

	//响应/query，从传入的参数删除一条记事本
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "failed request")
		} else {
			id := r.FormValue("id")
			fmt.Fprintln(w, query(id))
		}
	})

	//启动本地服务器（localhost）
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("server err：", err)
	}
}
