package probe

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

const (
	interval = 5 // 每隔5秒进行一次探测
)

func Host_probe() {
	// 要进行探测的IP地址和端口
	targets := map[string][]int{
		"127.0.0.1":   {3306, 6379},
		"192.168.0.2": {3306},
		"192.168.0.3": {6379},
	}

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

	// 开始进行探测
	for {
		for ip, ports := range targets {
			for _, port := range ports {
				address := fmt.Sprintf("%s:%d", ip, port)
				conn, err := net.DialTimeout("tcp", address, time.Second*1)
				if err != nil {
					log.WithFields(log.Fields{
						"Type":      "Host",
						"IP":        ip,
						"Port":      port,
						"err_reson": err,
					}).Error("主机 状态异常")
				} else {
					conn.Close()
					// log.Printf("%s:%d 可达\n", ip, port)
					log.WithFields(log.Fields{
						"Type": "Host",
						"IP":   ip,
						"Port": port,
					}).Info("主机 状态正常")
				}
			}
		}

		// 等待一段时间再进行下一次探测
		time.Sleep(time.Second * interval)
	}
}
