/**
 * @Author: entere
 * @Description:
 * @File:  unify
 * @Version: 1.0.0
 * @Date: 2020/3/22 13:21
 */

package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/entere/go-wechat-pay/mch/wxutils"

	"github.com/entere/go-wechat-pay/mch/core"
	"github.com/entere/go-wechat-pay/mch/pay"

	"github.com/gin-gonic/gin"
)

var (
	payment *pay.Pay
	outTradeNo = wxutils.TimeToString(time.Now())
)
func init() {
	Init()
	payment = pay.NewPay(core.NewClient(appID, mchID, apiKey, core.MD5), isSandbox)
}


// 下单 获取支付二维码
func UninfedOrder(c *gin.Context) {

	//// 获取沙箱api_key
	//sandboxApiKey, err := payment.GetSandboxSignKey(payment.WX_MCH_ID, payment.WX_API_KEY)
	//if err != nil {
	//	fmt.Printf("get sand box sign key err:%v\n", err)
	//}


	params := make(map[string]string)
	params["body"] = "test"
	params["out_trade_no"] = outTradeNo
	params["total_fee"] = strconv.FormatInt(int64(totalFree), 10)
	params["spbill_create_ip"] = "192.168.0.1"
	params["notify_url"] = notifyUrl
	params["trade_type"] = "NATIVE"

	resp, err := payment.UnifiedOrder(params)
	if err != nil {
		log.Fatalf("pay unifiedorder err:%v", err)

	}
	log.Printf("unified order result :%v\n", resp)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": resp,
	})

}


// 查询订单
func OrderQuery(c *gin.Context) {
	params := make(map[string]string)

	params["out_trade_no"] = outTradeNo
	params["nonce_str"] = wxutils.NonceStr(8)
	params["sign_type"] = core.MD5

	resp, err := payment.OrderQuery(params)
	if err != nil {
		log.Fatalf("pay unifiedorder err:%v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": resp,
	})

}


//支付通知
func Notify(c *gin.Context) {

	xmlByte, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("ioutil read error:%v", err)
	}

	resp, err := payment.Notify(string(xmlByte))
	if err != nil {
		fmt.Fprint(c.Writer, core.AckFail)
		log.Fatalf("pay unifiedorder err:%v", err)
	}
	//成功后给微信发ACK
	fmt.Fprint(c.Writer, core.AckSuccess)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": resp,
	})

}
