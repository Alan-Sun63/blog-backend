package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var db *gorm.DB

func init() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Printf("未找到 .env 文件")
	}
}

func main() {
	// 获取数据库连接字符串
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL 环境变量未设置")
	}

	// 连接数据库
	var err error
	db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	log.Println("连接数据库成功")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "孙将军牛逼",
		})
	})

	r.Run(":8080") // 默认监听 8080 端口
}
