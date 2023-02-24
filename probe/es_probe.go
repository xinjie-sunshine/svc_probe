package probe

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Es_probe() {
	// 创建日志文件
	file, err := os.OpenFile("probe.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)

	// 创建HTTP客户端
	httpClient := &http.Client{
		Timeout: time.Second * 5, // 设置5秒的超时时间
	}

	// 每1分钟探测一次Elasticsearch可用性
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		if isElasticsearchAvailable(httpClient) {
			log.Println("Elasticsearch 可用")
		} else {
			log.Println("Elasticsearch 不可用")
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
