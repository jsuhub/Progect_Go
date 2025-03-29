package main

import (
	"Progect/config"
	"log"
	"os"

	"Progect/router"
)

func init() {
	// 设置全局日志输出到文件
	f, err := os.OpenFile("/var/log/myapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	log.SetOutput(f)
}

func main() {
	log.Println("main: 启动程序")
	config.InitDB()
	r := router.SetRouter() // 带 Logger 和 Recovery 的引擎
	r.Run("0.0.0.0:8888")   // 默认监听 0.0.0.0:8888
}
