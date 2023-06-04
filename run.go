package main

import (
	"github.com/gin-gonic/gin"
	//可选
	_ "net/http"
)

func main() {
	//创建路由
	r := gin.Default()
	//绑定路由规则，执行的函数
	//gin.Context，封装了request和response
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "返回成功",
		})
	})
	//监听端口，默认在8080，可以在这里指定端口号
	r.Run(":8000")
	//r.Run()
}
