/**
 * @Author: entere
 * @Description:
 * @File:  constants
 * @Version: 1.0.0
 * @Date: 2020/3/23 21:47
 */

package core

const (
	Fail       = "FAIL"
	Success    = "SUCCESS"
	HMACSHA256 = "HMAC-SHA256"
	MD5        = "MD5"
	Sign       = "sign"

	MicroPayUrl         = "https://api.mch.weixin.qq.com/pay/micropay"
	UnifiedOrderUrl     = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	OrderQueryUrl       = "https://api.mch.weixin.qq.com/pay/orderquery"
	OpenidByCodeUrl     = "https://api.weixin.qq.com/sns/oauth2/access_token"
	AuthCodeToOpenidUrl = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"

	SandboxUnifiedOrderUrl = "https://api.mch.weixin.qq.com/sandboxnew/pay/unifiedorder"
	SandboxOrderQueryUrl   = "https://api.mch.weixin.qq.com/sandboxnew/pay/orderquery"

	SandboxGetSignKeyUrl       = "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"
	SandboxAuthCodeToOpenidUrl = "https://api.mch.weixin.qq.com/sandboxnew/tools/authcodetoopenid"

	AckSuccess = `<xml><return_code><![CDATA[SUCCESS]]></return_code></xml>`
	AckFail    = `<xml><return_code><![CDATA[FAIL]]></return_code></xml>`
)
