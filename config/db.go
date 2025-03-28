package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=gagamel  password=gagamel dbname=gagameldb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Println("数据库初始化成功", DB)
	if err != nil {
		panic(err)
	}

}
