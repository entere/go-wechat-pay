/**
 * @Author: entere
 * @Description:
 * @File:  unify
 * @Version: 1.0.0
 * @Date: 2020/3/22 13:21
 */

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/entere/go-wechat-pay/mch/wxutils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/entere/go-wechat-pay/mch/core"
	"github.com/entere/go-wechat-pay/mch/pay"

	"github.com/gin-gonic/gin"
)

var (
	payment *pay.Pay
)

func init() {
	Init()
	payment = pay.NewPay(core.NewCore(WXAppID, WXMchID, WXApiKey, core.MD5), WXIsSandbox)

}

// 下单 获取支付二维码
func UninfedOrderNative(c *gin.Context) {

	//// 获取沙箱api_key
	//sandboxApiKey, err := payment.GetSandboxSignKey(payment.WX_MCH_ID, payment.WX_API_KEY)
	//if err != nil {
	//	fmt.Printf("get sand box sign key err:%v\n", err)
	//}
	outTradeNo := wxutils.TimeToString(time.Now())
	log.Printf("unified order out trade no :%v\n", outTradeNo)

	params := make(map[string]string)
	params["body"] = "test"
	params["out_trade_no"] = outTradeNo
	params["total_fee"] = strconv.FormatInt(int64(WXTotalFree), 10)
	params["spbill_create_ip"] = "192.168.0.1"
	params["notify_url"] = WXNotifyUrl
	params["trade_type"] = "NATIVE"

	resp, err := payment.UnifiedOrder(params)

	if err != nil {
		log.Fatalf("pay unifiedorder err:%v", err)
	}

	log.Printf("unified order result :%v\n", resp)
	resp["out_trade_no"] = outTradeNo
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": resp,
	})

}

// JsAPI统一下单接口
func UninfedOrderJSAPI(c *gin.Context) {

	//// 获取沙箱api_key
	//sandboxApiKey, err := payment.GetSandboxSignKey(payment.WX_MCH_ID, payment.WX_API_KEY)
	//if err != nil {
	//	fmt.Printf("get sand box sign key err:%v\n", err)
	//}
	openid := c.PostForm("openid")
	fmt.Printf("openid is aaa:%v\n", openid)
	outTradeNo := wxutils.TimeToString(time.Now())

	params := make(map[string]string)
	params["body"] = "test"
	params["out_trade_no"] = outTradeNo
	params["total_fee"] = strconv.FormatInt(int64(WXTotalFree), 10)
	params["spbill_create_ip"] = "192.168.0.1"
	params["notify_url"] = WXNotifyUrl
	params["trade_type"] = "JSAPI"
	params["openid"] = openid

	fmt.Printf("params is  :%v\n", params)

	resp, err := payment.UnifiedOrder(params)
	resp["out_trade_no"] = outTradeNo
	if err != nil {
		log.Fatalf("pay unifiedorder err:%v", err)
	}
	log.Printf("unified order result :%v\n", resp)

	jsParams := make(map[string]string)
	jsParams["appId"] = WXAppID
	jsParams["timeStamp"] = wxutils.TimeToString(time.Now())
	jsParams["nonceStr"] = wxutils.NonceStr(10)
	jsParams["package"] = "prepay_id=" + resp["prepay_id"]
	jsParams["signType"] = core.MD5
	paySign := wxutils.Sign(jsParams, WXApiKey, jsParams["signType"])
	jsParams["paySign"] = paySign

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": jsParams,
	})

}

// 查询订单
func OrderQuery(c *gin.Context) {
	outTradeNo := c.PostForm("outTradeNo")
	log.Printf("order query out trade no :%v\n", outTradeNo)
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

//Native支付回调通知
func NotifyNative(c *gin.Context) {

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

// 跳转获取openid
func LoginCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		redirectURI := "http://pay.raccooncode.com/weixin/login/mp/callback"
		redirectURI = url.QueryEscape(redirectURI)
		wxURL := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + WXAppID + "&redirect_uri=" + redirectURI + "&response_type=code&scope=snsapi_base&state=123#wechat_redirect"

		c.Redirect(http.StatusFound, wxURL)
	}
	log.Printf("code is:%v", code)

	h := &http.Client{}
	response, err := h.Get("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + WXAppID + "&secret=" + WXAppSecret + "&code=" + code + "&grant_type=authorization_code")
	if err != nil {
		fmt.Errorf("get openid err:%v", err)
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)

	var mapResult map[string]interface{}
	err = json.Unmarshal(res, &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fmt.Println(mapResult)

	if err != nil {
		fmt.Printf("get openid ioutil read err:%v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": string(res),
	})

}

func GetParams(c *gin.Context) {

	code := c.PostForm("code")

	fmt.Printf("aaa: %v", code)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": string(code),
	})
}
