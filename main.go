package main

// 导入 Gin 框架
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的 Gin 路由器，包括 Logger 和 Recovery 中间件。
	r := gin.Default()

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
