package probe

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Mysql_probe_test() {
	// Load configuration from file
	viper.Reset()
	viper.SetConfigName("mysql_cnf")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to read configuration file")
	}

	// Initialize logger
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// logrus.SetOutput(os.Stdout)
	logFile, err := os.OpenFile("probe.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to open log file")
	}
	logrus.SetOutput(logFile)
	logrus.SetLevel(logrus.InfoLevel)

	dbHost := viper.GetString("dbHost")
	dbPort := viper.GetString("dbPort")
	dbUser := viper.GetString("dbUser")
	dbPassword := viper.GetString("dbPassword")
	dbName := viper.GetString("dbName")
	// fmt.Print(dbHost, dbName)

	// 连接 MySQL 数据库
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		logrus.Printf("连接 MySQL 数据库失败：%v", err)
	}

	// 检查 MySQL 数据库的可用性
	if err = db.Ping(); err != nil {
		logrus.WithFields(logrus.Fields{
			"Type":      "MySQl",
			"IP":        dbHost,
			"Port":      dbPort,
			"Err_reson": err,
		}).Error("MySql 状态异常")
	} else {
		logrus.WithFields(logrus.Fields{
			"Type": "MySQl",
			"IP":   dbHost,
			"Port": dbPort,
		}).Info("MySql 状态正常")
	}

	db.Close()
}
