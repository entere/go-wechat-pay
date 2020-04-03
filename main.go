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

	// Native支付
	r.POST("/wechat/pay/unifiedorder/native", handler.UninfedOrderNative)
	r.POST("/wechat/pay/notify/native", handler.NotifyNative)
	// JSAPI支付
	r.POST("/wechat/pay/unifiedorder/jsapi", handler.UninfedOrderJSAPI)
	//微信支付回调方法

	//微信支付查询方法
	r.POST("/wechat/pay/orderquery", handler.OrderQuery)

	r.GET("/wechat/pay/openid", handler.GetOpenID)

	//test
	r.POST("/wechat/pay/getparams", handler.GetParams)

	r.Run(":8003") // 监听并在 0.0.0.0:8080 上启动服务
}
