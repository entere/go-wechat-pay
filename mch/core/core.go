/**
 * @Author: entere
 * @Description:
 * @File:  Core
 * @Version: 1.0.0
 * @Date: 2020/3/23 22:30
 */

package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/entere/go-wechat-pay/mch/wxutils"
)

type Core struct {
	appID    string
	mchID    string
	apiKey   string
	signType string
}

func NewCore(appID, mchID, apiKey, signType string) *Core {

	return &Core{
		appID:    appID,
		mchID:    mchID,
		apiKey:   apiKey,
		signType: signType,
	}
}



// 微信支付签名，微信文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_3
func (c *Core) Sign(params map[string]string) string {
	// 创建切片
	var keys = make([]string, 0, len(params))
	// 遍历签名参数
	for k := range params {
		if k != "sign" { // 排除sign字段
			keys = append(keys, k)
		}
	}

	// 由于切片的元素顺序是不固定，所以这里强制给切片元素加个顺序
	sort.Strings(keys)

	//创建字符缓冲
	var buf bytes.Buffer
	for _, k := range keys {
		if len(params[k]) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params[k])
			buf.WriteString(`&`)
		}
	}
	// 加入apiKey作加密密钥
	buf.WriteString(`key=`)
	buf.WriteString(c.apiKey)

	var (
		dataMd5    [16]byte
		dataSha256 []byte
		str        string
	)

	switch c.signType {
	case MD5:
		dataMd5 = md5.Sum(buf.Bytes())
		str = hex.EncodeToString(dataMd5[:]) //需转换成切片
	case HMACSHA256:
		h := hmac.New(sha256.New, []byte(c.apiKey))
		h.Write(buf.Bytes())
		dataSha256 = h.Sum(nil)
		str = hex.EncodeToString(dataSha256[:])
	}

	return strings.ToUpper(str)
}

// 生成带有签名的xml字符串
func (c *Core) GenerateSignedXML(params map[string]string) string {
	sign := c.Sign(params)
	params[Sign] = sign
	return wxutils.Map2XML(params)
}

// post请求，without fillRequestData
func (c *Core) PostWithoutFillRequestData(url string, params map[string]string) (string, error) {
	h := &http.Client{}
	params["sign"] = c.Sign(params)
	response, err := h.Post(url, "application/xml; charset=utf-8", strings.NewReader(wxutils.Map2XML(params)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// post请求，without cert
func (c *Core) PostWithoutCert(url string, params map[string]string) (string, error) {
	h := &http.Client{}
	p := c.fillRequestData(params)
	response, err := h.Post(url, "application/xml; charset=utf-8", strings.NewReader(wxutils.Map2XML(p)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// 验证签名
func (c *Core) ValidSign(params map[string]string) bool {
	if _, ok := params[Sign]; !ok {
		return false
	}
	return params[Sign] == c.Sign(params)
}

// 向 params 中添加 appid、mch_id、nonce_str、sign_type、sign
func (c *Core) fillRequestData(params map[string]string) map[string]string {
	params["appid"] = c.appID
	params["mch_id"] = c.mchID
	params["nonce_str"] = wxutils.NonceStr(8)
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)
	return params
}

// 处理 HTTPS API返回数据，转换成Map对象。return_code为SUCCESS时，验证签名。
func (c *Core) ProcessResponseXML(xmlStr string) (map[string]string, error) {
	var returnCode string
	params, err := wxutils.XML2Map(xmlStr)
	if err != nil {
		return nil, err

	}
	if _, ok := params["return_code"]; ok {
		returnCode = params["return_code"]
	} else {
		return nil, errors.New("no return_code in XML")
	}
	if returnCode == Fail {
		return params, nil
	} else if returnCode == Success {
		if c.ValidSign(params) {
			return params, nil
		} else {
			return nil, errors.New("invalid sign value in XML")
		}
	} else {
		return nil, errors.New("return_code value is invalid in XML")
	}
}
