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

	// 每分钟执行一次探测
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		probe.Host_probe_test()
		probe.Es_probe()
		// probe.Api_probe_test()
		// probe.Redis_probe_test()
		// probe.Mysql_probe_test()
	}

	// Wait for signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals

	// Stop probe and exit
	log.Print("Stopping probe")
	os.Exit(0)

}
