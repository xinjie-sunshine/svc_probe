package probe

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func Api_probe() {

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

	// 多个API地址
	apiAddrs := []string{"https://www.baidu1.com", "https://cn.bing.com/"}

	// // 设置日志输出到文件
	// log.SetOutput(file)
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		for _, apiAddr := range apiAddrs {
			// 发送HTTP HEAD请求探测API存活状态
			resp, err := http.Head(apiAddr)

			if err != nil || resp.StatusCode != http.StatusOK {
				// API不可用
				log.WithFields(log.Fields{
					"url":    apiAddr,
					"status": 500,
				}).Error("API 状态码异常")
			} else {
				// API可用
				log.WithFields(log.Fields{
					"url":    apiAddr,
					"status": 200,
				}).Info("API 状态码正常")
			}
		}
	}
}
