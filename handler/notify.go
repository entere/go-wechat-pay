/**
 * @Author: entere
 * @Description:
 * @File:  unify
 * @Version: 1.0.0
 * @Date: 2020/3/22 13:21
 */

package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/entere/go-wechat-pay/mch/core"
	"github.com/entere/go-wechat-pay/mch/pay"

	"github.com/gin-gonic/gin"
)

func Notify(c *gin.Context) {

	xmlByte, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("ioutil read error:%v", err)
	}

	pay := pay.NewPay(core.NewClient(core.WX_APP_ID, core.WX_MCH_ID, core.WX_API_KEY_SANBOX, core.MD5))
	resp, err := pay.Notify(string(xmlByte))
	if err != nil {
		log.Fatalf("pay unifiedorder err:%v", err)

	}
	log.Printf("notify result :%v\n", resp)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": resp,
	})

}
