package probe

import (
	"fmt"
	"log"
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
	logFile, err := os.OpenFile("probe.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("无法创建log文件：", err)
		return
	}

	// 设置日志输出到文件
	log.SetOutput(logFile)
	defer logFile.Close()

	// 开始进行探测
	for {
		for ip, ports := range targets {
			for _, port := range ports {
				address := fmt.Sprintf("%s:%d", ip, port)
				conn, err := net.DialTimeout("tcp", address, time.Second*1)
				if err != nil {
					log.Printf("%s:%d 不可达,原因为%s\n", ip, port, err)
				} else {
					conn.Close()
					log.Printf("%s:%d 可达\n", ip, port)
				}
			}
		}

		// 等待一段时间再进行下一次探测
		time.Sleep(time.Second * interval)
	}
}
