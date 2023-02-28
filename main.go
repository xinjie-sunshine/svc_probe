package main

import (
	. "svc_probe/probe"
)

func main() {
	go Api_probe()
	go Mysql_probe()
	go Es_probe()
	go Redis_probe()
	Host_probe()

}
