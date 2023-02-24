package probe

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbName     = "mydb"
	dbUser     = "root"
	dbPassword = "root"
	dbHost     = "localhost"
	dbPort     = "3306"
)

func Mysql_probe() {

	// 创建日志文件
	file, err := os.OpenFile("probe.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)

	// 设置定时任务每隔一分钟执行一次
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// 连接 MySQL 数据库
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
		if err != nil {
			log.Printf("连接 MySQL 数据库失败：%v", err)
			continue
		}

		// 检查 MySQL 数据库的可用性
		if err = db.Ping(); err != nil {
			log.Printf("MySQL 服务不可用：%v", err)
		} else {
			log.Println("MySQL 服务正常")
		}

		db.Close()
	}
}
