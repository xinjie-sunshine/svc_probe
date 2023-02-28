package probe

import (
	"context"
	log "github.com/sirupsen/logrus"
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
			log.WithFields(log.Fields{
				"Type":      "Redis",
				"IP":        redisHost,
				"Port":      redisPort,
				"Err_reson": err,
			}).Error("Redis 状态异常")

		} else {
			log.WithFields(log.Fields{
				"Type": "Redis",
				"IP":   redisHost,
				"Port": redisPort,
			}).Info("Redis 状态异常")
		}

	}
}
