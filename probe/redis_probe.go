package probe

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func Redis_probe_test() {
	// Load configuration from file
	viper.Reset()
	viper.SetConfigName("redis_cnf")
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

	//从配置文件获取-未获取到数据
	addr_and_ports := viper.GetString("addr_and_ports")

	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     addr_and_ports,
		Password: "",
		DB:       0,
	})

	// 探测Redis存活状态
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Type":      "Redis",
			"host":      addr_and_ports,
			"Err_reson": err,
		}).Error("Redis 状态异常")
	} else {
		logrus.WithFields(logrus.Fields{
			"Type": "Redis",
			"host": addr_and_ports,
		}).Info("Redis 状态异常")
	}

}
