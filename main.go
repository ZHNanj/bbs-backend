package main

// 导入 Gin 框架
import (
	"bbs-backend/config"
	middlewares "bbs-backend/middleware"
	"bbs-backend/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/http"
)

func main() {
	// 连接数据库
	cfg := config.New()
	repository.InitDB(cfg)

	// 配置数据库版本控制
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	// 创建一个默认的 Gin 路由器，包括 Logger 和 Recovery 中间件。
	r := gin.Default()

	// 配置跨域
	r.Use(middlewares.Cors())

	// 定义路由规则，当 HTTP GET 请求匹配到 URL 路径 /health 时。执行作为参数的请求函数。gin.Context 参数封装了 HTTP 请求的所有详细信息和构造响应的方法。
	r.GET("/health", func(context *gin.Context) {
		// 发送 JSON 格式的响应。gin.H{...} 用于创建一个 map，以构建 JSON 对象。
		context.JSON(http.StatusOK, gin.H{
			"message": "hello, gin!",
			"data":    "请求成功！",
		})
	})
	// 默认 8080 端口启动 HTTP 服务器
	r.Run()

	// 指定端口启动 HTTP 服务器
	// r.r.Run(":8888")
}
