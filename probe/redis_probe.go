package probe

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	// Redis连接信息
	redisHost     = "localhost"
	redisPort     = "6379"
	redisPassword = ""
)

func Redis_probe() {

	// 创建日志文件
	file, err := os.OpenFile("probe.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)

	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	// 每分钟执行一次探测
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		// 探测Redis存活状态
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Printf("redis_client异常", err)
		} else {
			log.Println("redis_client服务正常")
		}

	}
}
