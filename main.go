/**
 * @Author: entere
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/3/21 09:35
 */

package main

import (
	"github.com/entere/go-wechat-pay/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//微信支付回调方法
	r.POST("/wechat/pay/unifiedorder", handler.UninfedOrder)
	r.POST("/wechat/pay/notify", handler.Notify)

	r.Run(":8003") // 监听并在 0.0.0.0:8080 上启动服务
}
