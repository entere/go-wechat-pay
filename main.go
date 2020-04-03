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

	// Native扫码支付统一单
	r.POST("/pay/unifiedorder/native", handler.UninfedOrderNative)
	// Native扫码支付成功通知
	r.POST("/pay/notify/native", handler.NotifyNative)
	// Native扫码支付查询订单
	r.POST("/pay/orderquery", handler.OrderQuery)

	// JSAPI支付统一下单
	r.POST("/pay/unifiedorder/jsapi", handler.UninfedOrderJSAPI)
	// JSAPI支付获取openid
	r.GET("/pay/openid", handler.GetOpenID)

	//微信支付查询方法

	//test
	r.POST("/pay/getparams", handler.GetParams)

	r.Run(":8003") // 监听并在 0.0.0.0:8080 上启动服务
}
