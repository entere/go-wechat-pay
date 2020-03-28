/**
 * @Author: entere
 * @Description:
 * @File:  handler
 * @Version: 1.0.0
 * @Date: 2020/3/25 22:04
 *
 */

package handler

var WXIsSandbox = true

var (
	WXAppID     string
	WXAppSecret string
	WXMchID     string
	WXApiKey    string
	WXTotalFree int64
	WXNotifyUrl string
)


func Init() {

	WXAppID     = "xx"
	WXAppSecret = "xx"
	WXMchID     = "xx"
	WXApiKey    = "xx"




	//todo 不用沙箱  换成配置文件
	if !WXIsSandbox {
		WXTotalFree = 1
		WXNotifyUrl = "http://api.raccooncode.com/go-wechat-pay/wechat/pay/notify"
	} else {
		WXApiKey    = GetSandboxApiKey(WXAppID,WXMchID,WXApiKey)
		WXTotalFree = 301
		WXNotifyUrl = "http://6daf869b.ngrok.io/wechat/pay/notify"
	}


}

