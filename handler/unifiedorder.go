/**
 * @Author: entere
 * @Description:
 * @File:  unify
 * @Version: 1.0.0
 * @Date: 2020/3/22 13:21
 */

package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/entere/go-wechat-pay/mch/wxutils"

	"github.com/entere/go-wechat-pay/mch/core"
	"github.com/entere/go-wechat-pay/mch/pay"

	"github.com/gin-gonic/gin"
)

const AssetsUrl = "http://wxpayment.free.idcfengye.com/assets"

func UninfedOrder(c *gin.Context) {

	//// 获取沙箱api_key
	//sandboxApiKey, err := payment.GetSandboxSignKey(payment.WX_MCH_ID, payment.WX_API_KEY)
	//if err != nil {
	//	fmt.Printf("get sand box sign key err:%v\n", err)
	//}

	params := make(map[string]string)
	params["body"] = "test"
	params["out_trade_no"] = wxutils.TimeToString(time.Now())
	params["total_fee"] = strconv.FormatInt(301, 10)
	params["spbill_create_ip"] = "192.168.0.1"
	params["notify_url"] = "http://6825eec0.ngrok.io/wechat/pay/notify"
	params["trade_type"] = "NATIVE"
	pay := pay.NewPay(core.NewClient(core.WX_APP_ID, core.WX_MCH_ID, core.WX_API_KEY_SANBOX, core.MD5))
	resp, err := pay.UnifiedOrder(params)
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
