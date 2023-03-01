package probe

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Api_probe_test() {
	// Load configuration from file
	viper.Reset()
	viper.SetConfigName("api_cnf")
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
	// log.Print(hostsAndPorts)
	// timeout := time.Duration(viper.GetInt("timeout")) * time.Second
	// timeout := time.Duration(5) * time.Second

	for _, apiAddr := range hostsAndPorts {
		// 发送HTTP HEAD请求探测API存活状态
		resp, err := http.Head(apiAddr)

		if err != nil || resp.StatusCode != http.StatusOK {
			// API不可用
			logrus.WithFields(logrus.Fields{
				"url":    apiAddr,
				"status": 500,
			}).Error("API 状态码异常")
		} else {
			// API可用
			logrus.WithFields(logrus.Fields{
				"url":    apiAddr,
				"status": 200,
			}).Info("API 状态码正常")
		}
	}
}
