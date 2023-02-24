package probe

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Api_probe() {
	// 多个API地址
	apiAddrs := []string{"https://www.baidu.com", "https://pay.sc.189.cn"}

	// 创建日志文件
	file, err := os.OpenFile("probe.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to create log file: %v", err)
	}
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		for _, apiAddr := range apiAddrs {
			// 发送HTTP HEAD请求探测API存活状态
			resp, err := http.Head(apiAddr)

			if err != nil || resp.StatusCode != http.StatusOK {
				// API不可用
				log.Printf("API %s 状态码异常 \n", apiAddr)
			} else {
				// API可用
				log.Printf("API %s 状态码正常 \n", apiAddr)
			}
		}
	}
}
