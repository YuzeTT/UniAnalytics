package util

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func createSql() {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	table := `
		CREATE TABLE "links" (
			"uid"	INTEGER NOT NULL UNIQUE,
			"created_at"	TEXT NOT NULL,
			"url"	TEXT NOT NULL,
			"tip_page"	BLOB NOT NULL DEFAULT 'false',
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
