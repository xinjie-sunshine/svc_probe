package probe

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func Redis_probe() {
	// Redis连接信息
	redisHost := "localhost"
	redisPort := "6379"
	redisPassword := ""

	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	// 每分钟执行一次探测
	ticker := time.NewTicker(1 * time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		// 探测Redis存活状态
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("redis_client异常", err)
		}

		// 将探测结果写入日志文件
		f, err := os.OpenFile("probe.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("无法打开probe.log文件: %v", err)
		}
		defer f.Close()

		logger := log.New(f, "", log.LstdFlags)
		if err != nil {
			logger.Printf("Redis节点探测失败: %v", err)
		} else {
			logger.Printf("Redis节点存活状态: 存活")
		}
	}
}
