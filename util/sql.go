package util

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func createSql() {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	table := `
		CREATE TABLE "links" (
			"uid"	INTEGER NOT NULL UNIQUE,
			"created_at"	TEXT NOT NULL,
			"url"	TEXT NOT NULL,
			"ip"	TEXT,
			"type"	INTEGER NOT NULL DEFAULT 0,
			PRIMARY KEY("uid" AUTOINCREMENT)
		);
	`

	_, err = db.Exec(table)
	checkErr(err)
	if err == nil {
		log.Println("数据库已创建")
	}

	db.Close()
}

func InitSql() {
	log.Println("开始初始化数据库")
	isFile := FileExist("./data.db")
	if !isFile {
		log.Println("未找到数据库文件，开始自动创建")
		createSql()
	}
	log.Println("数据库初始化完毕")
}

func AddSql(url string, ip string) {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	stmt, err := db.Prepare("INSERT INTO links(uid, created_at, url, ip, type) values(?,?,?,?,?)")
	checkErr(err)
	stmt.Exec(nil, time.Now().Format("2006-01-02 15:04:05"), url, ip, 0)
	// res, err := stmt.Exec(nil, time.Now().Format("2006-01-02 15:04:05"), url, 0)
	// checkErr(err)
	// id, err := res.LastInsertId()
	// checkErr(err)
	// fmt.Println(id)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
