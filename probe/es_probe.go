package probe

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Es_probe() {
	// Load configuration from file
	viper.Reset()
	viper.SetConfigName("es_cnf")
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

	hostsAndPorts := viper.GetStringSlice("addr_and_ports")
	// 创建HTTP客户端
	httpClient := &http.Client{
		Timeout: time.Second * 5, // 设置5秒的超时时间
	}

	for _, apiAddr := range hostsAndPorts {
		// 发送HTTP HEAD请求探测API存活状态
		resp, err := httpClient.Get(apiAddr)
		if err != nil || resp.StatusCode != http.StatusOK {
			// API不可用
			logrus.WithFields(logrus.Fields{
				"url":    apiAddr,
				"status": 500,
			}).Error("ES 状态异常")
		} else {
			// API可用
			logrus.WithFields(logrus.Fields{
				"url":    apiAddr,
				"status": 200,
			}).Info("ES 状态正常")
		}

	}
}
