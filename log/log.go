package log

import (
	"log"
	"os"
)

func Start() {
	f, err := os.OpenFile("/var/log/myapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("服务启动，日志开始记录")
}
