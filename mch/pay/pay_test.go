/**
 * @Author: entere
 * @Description:
 * @File:  pay_test
 * @Version: 1.0.0
 * @Date: 2020/3/24 22:31
 */

package pay

import (
	"strconv"
	"testing"
	"time"

	"github.com/entere/go-wechat-pay/mch/wxutils"

	"github.com/entere/go-wechat-pay/mch/core"
)

var (
	testAppID     = "xx"
	testAppSecret = "xx"
	testMchID     = "xx"
	testApiKey    = "xx" // sanbox api key 需能过接口获取
)

func TestPay_UnifiedOrder(t *testing.T) {

	params := make(map[string]string)
	params["body"] = "test"
	params["out_trade_no"] = wxutils.TimeToString(time.Now())
	params["total_fee"] = strconv.FormatInt(301, 10)
	params["spbill_create_ip"] = "192.168.0.1"
	params["notify_url"] = "http://localhost:8003/wechat/pay/notify"
	params["trade_type"] = "NATIVE"
	pay := NewPay(core.NewCore(testAppID, testMchID, testApiKey, core.MD5), true)
	resp, err := pay.UnifiedOrder(params)
	if err != nil {
		t.Fatalf("pay unifiedorder err:%v", err)

	}
	t.Logf("unified order result :%v", resp)
}

func TestPay_OrderQuery(t *testing.T) {
	pay := NewPay(core.NewCore(testAppID, testMchID, testApiKey, core.MD5), true)
	params := make(map[string]string)
	params["out_trade_no"] = "202008089877"
	params["nonce_str"] = wxutils.NonceStr(8)
	params["sign_type"] = core.MD5

	resp, err := pay.OrderQuery(params)
	if err != nil {
		t.Fatalf("pay unifiedorder err:%v", err)

	}
	t.Logf("order query result :%v", resp)
}
