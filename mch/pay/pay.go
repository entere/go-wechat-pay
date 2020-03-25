/**
 * @Author: entere
 * @Description:
 * @File:  pay
 * @Version: 1.0.0
 * @Date: 2020/3/24 15:03
 */

package pay

import (
	"log"

	"github.com/entere/go-wechat-pay/mch/core"
)

type Pay struct {
	cli *core.Client
}

func NewPay(cli *core.Client) *Pay {
	return &Pay{
		cli: cli,
	}
}

// 统一下单
func (p *Pay) UnifiedOrder(params map[string]string) (map[string]string, error) {
	var url string
	if core.IsSandbox {
		url = core.SandboxUnifiedOrderUrl
	} else {
		url = core.UnifiedOrderUrl
	}
	xmlStr, err := p.cli.PostWithoutCert(url, params)
	if err != nil {
		return nil, err
	}
	return p.cli.ProcessResponseXml(xmlStr)
}

// 订单查询
func (p *Pay) OrderQuery(params map[string]string) (map[string]string, error) {
	var url string
	if core.IsSandbox {
		url = core.SandboxOrderQueryUrl
	} else {
		url = core.OrderQueryUrl
	}
	xmlStr, err := p.cli.PostWithoutCert(url, params)
	if err != nil {
		return nil, err
	}
	return p.cli.ProcessResponseXml(xmlStr)
}

//支付回调接收通知接口
func (p *Pay) Notify(xmlStr string) (map[string]string, error) {
	resp, err := p.cli.ProcessResponseXml(xmlStr)

	log.Printf("reveive notify:%+v", resp)
	if err != nil {

		return nil, err
	}

	return resp, nil
}
