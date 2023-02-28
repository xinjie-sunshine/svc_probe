package probe

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func Es_probe() {
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

	// 创建HTTP客户端
	httpClient := &http.Client{
		Timeout: time.Second * 5, // 设置5秒的超时时间
	}

	// 每1分钟探测一次Elasticsearch可用性
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		if isElasticsearchAvailable(httpClient) {
			log.WithFields(log.Fields{
				"type":   "elasticSearch",
				"status": 200,
			}).Info("elasticSearch 状态正常")
		} else {
			log.WithFields(log.Fields{
				"type":   "elasticSearch",
				"status": 500,
			}).Error("elasticSearch 状态异常")
		}
	}
}

func isElasticsearchAvailable(httpClient *http.Client) bool {
	resp, err := httpClient.Get("http://localhost:9200/_cluster/health")
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
