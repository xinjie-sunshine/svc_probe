package probe

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
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
	file, err := os.OpenFile("probe.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to create log file: %v", err)
	}
	defer file.Close()
	// 日志作为JSON而不是默认的ASCII格式器.
	log.SetFormatter(&log.JSONFormatter{})

	// 输出到标准输出,可以是任何io.Writer
	log.SetOutput(file)

	// 只记录xx级别或以上的日志
	log.SetLevel(log.TraceLevel)

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
			log.WithFields(log.Fields{
				"Type":      "MySQl",
				"IP":        dbHost,
				"Port":      dbPort,
				"Err_reson": err,
			}).Error("MySql 状态异常")
		} else {
			log.WithFields(log.Fields{
				"Type": "MySQl",
				"IP":   dbHost,
				"Port": dbPort,
			}).Info("主机 状态正常")
		}

		db.Close()
	}
}
