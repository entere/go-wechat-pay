/**
 * @Author: entere
 * @Description:
 * @File:  pay
 * @Version: 1.0.0
 * @Date: 2020/3/24 15:03
 */

package pay

import (
	"encoding/json"
	"fmt"
	"github.com/entere/go-wechat-pay/mch/core"
	"io/ioutil"
	"log"
	"net/http"
)

type Pay struct {
	cli       *core.Core
	isSandbox bool
}

func NewPay(cli *core.Core, isSandbox bool) *Pay {
	return &Pay{
		cli:       cli,
		isSandbox: isSandbox,
	}
}

// 统一下单
func (p *Pay) UnifiedOrder(params map[string]string) (map[string]string, error) {
	var url string
	if p.isSandbox {
		url = core.SandboxUnifiedOrderUrl
	} else {
		url = core.UnifiedOrderUrl
	}
	xmlStr, err := p.cli.PostWithoutCert(url, params)
	if err != nil {
		return nil, err
	}
	return p.cli.ProcessResponseXML(xmlStr)
}

// 订单查询
func (p *Pay) OrderQuery(params map[string]string) (map[string]string, error) {
	var url string
	if p.isSandbox {
		url = core.SandboxOrderQueryUrl
	} else {
		url = core.OrderQueryUrl
	}
	xmlStr, err := p.cli.PostWithoutCert(url, params)
	if err != nil {
		return nil, err
	}
	return p.cli.ProcessResponseXML(xmlStr)
}

//支付回调接收通知接口
func (p *Pay) Notify(xmlStr string) (map[string]string, error) {
	resp, err := p.cli.ProcessResponseXML(xmlStr)

	log.Printf("reveive notify:%+v", resp)
	if err != nil {

		return nil, err
	}

	return resp, nil
}

// 授权码查询OPENID接口

func (p *Pay) AuthCodeToOpenid(params map[string]string) (map[string]string, error) {
	var url string
	if p.isSandbox {
		url = core.SandboxAuthCodeToOpenidUrl
	} else {
		url = core.AuthCodeToOpenidUrl
	}
	xmlStr, err := p.cli.PostWithoutCert(url, params)
	if err != nil {
		return nil, err
	}
	return p.cli.ProcessResponseXML(xmlStr)
}

func (p *Pay) GetOpenidByCode(code string, appID string, appSecret string) (map[string]interface{}, error) {
	var url string
	url = core.OpenidByCodeUrl
	h := &http.Client{}
	response, err := h.Get(url + "?appid=" + appID + "&secret=" + appSecret + "&code=" + code + "&grant_type=authorization_code")
	if err != nil {
		fmt.Printf("get openid err:%v", err)
		return nil, err

	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var ret map[string]interface{}
	if err := json.Unmarshal(res, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}
