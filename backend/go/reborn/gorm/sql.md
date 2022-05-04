
### database/sql
Go 语言标准库中 database/sql 包，用于查询各种 SQL 数据库。它将所有通用 SQL 功能抽象到一个 API中供开发者使用。但是 Go 的标准库中不包括数据库驱动程序。数据库驱动程序由特定软件包提供的，用于实现特定数据库底层的封装。这对于向前兼容很有用，也使得 Go 不会变得臃肿。因为在创建所有 Go 软件包时，开发人员无法预见未来会有什么数据库会被投入使用，而且要支持每个可能的数据库将需要进行大量维护工作。  

[更多可参考](http://go-database-sql.org/index.html)
```go
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func DbPing() {
	// 获得数据库对象，这里打开数据连接
	db, err := sql.Open("mysql", "demo:demo@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping的时候会与数据库建立连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected")
}

func MysqlDemoCode() {
	db, err := sql.Open("mysql", "demo:demo@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // 创建表
		query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // 插入新数据
		username := "Joshua"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ // 查询单个用户
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // 查询所有用户
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
```