package main

import (
	"Progect/config"
	"Progect/router"
)

func main() {
	config.InitDB()
	r := router.SetRouter() // 带 Logger 和 Recovery 的引擎
	r.Run("0.0.0.0:8888")   // 默认监听 0.0.0.0:8080
}
