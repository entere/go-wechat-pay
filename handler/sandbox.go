/**
 * @Author: entere
 * @Description:
 * @File:  sandbox
 * @Version: 1.0.0
 * @Date: 2020/3/27 23:37
 */

package handler

import (
    "github.com/entere/go-wechat-pay/mch/core"
    "github.com/entere/go-wechat-pay/mch/wxutils"
    "log"
)


// 获取沙箱验证签名的密钥key，沙箱环境用此密钥，正式环境还是用商户后台的api_key，
func GetSandboxApiKey(appID string,mchID string,apiKey string) string{
    c:=core.NewCore(appID, mchID, apiKey, core.MD5)
    params := make(map[string]string)
    params["mch_id"] = mchID
    params["nonce_str"] = wxutils.NonceStr(8)

    respXML,err:=c.PostWithoutFillRequestData(core.SandboxGetSignKeyUrl,params)
    if err != nil {
        log.Fatalf("get sandbox api_key PostWithoutCert err:%v",err)
    }

    resp, err := wxutils.XML2Map(respXML)
    if err != nil {
        log.Fatalf("get sandbox api_key ProcessResponseXml err:%v",err)
    }
    return resp["sandbox_signkey"]


}