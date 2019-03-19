package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db 是一个 sql.DB 类型的对象
	// 该对象线程安全，且内部已包含了一个连接池
	// 连接池的选项可以在 sql.Open 中设置，这里为了简单省略了
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	rows, err := db.Query("select username from user where user_id = ?", 2)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// 必须要把 rows 里的内容读完，否则连接永远不会释放
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
