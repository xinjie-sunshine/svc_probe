package probe

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Mysql_probe() {
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

	con_stings := viper.GetStringSlice("hosts_and_ports")
	for _, hostAndPort := range con_stings {
		// 连接 MySQL 数据库
		db, err := sql.Open("mysql", hostAndPort)
		host := strings.Split(hostAndPort, "@")[1]
		if err != nil {
			logrus.Printf("连接 MySQL 数据库失败：%v", err)
		}

		defer db.Close()
		// 检查 MySQL 数据库的可用性
		if err = db.Ping(); err != nil {
			logrus.WithFields(logrus.Fields{
				"Type":      "MySQl",
				"Host":      host,
				"Err_reson": err,
			}).Error("MySql 状态异常")
		} else {
			logrus.WithFields(logrus.Fields{
				"Type": "MySQl",
				"Host": host,
			}).Info("MySql 状态正常")
		}

	}
}
