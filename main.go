package main

import (
	"log"
	"os"
	"os/signal"
	"svc_probe/probe"
	"syscall"
	"time"
)

func main() {
	// a := "root:root@tcp(localhost:3307)/mydb"
	// b := strings.Split("@", a)
	// fmt.Print(b)

	// 每分钟执行一次探测
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		probe.Host_probe()
		probe.Es_probe()
		probe.Api_probe()
		probe.Redis_probe()
		probe.Mysql_probe()
	}

	// Wait for signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals

	// Stop probe and exit
	log.Print("Stopping probe")
	os.Exit(0)

}
