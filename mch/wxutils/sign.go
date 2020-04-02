/**
 * @Author: entere
 * @Description:
 * @File:  sign
 * @Version: 1.0.0
 * @Date: 2020/4/2 09:28
 */

package wxutils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

// @description 微信支付签名，微信文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_3
// @param params 需要签名的参数
// @return 返回签名后的字符串

func Sign(params map[string]string, apiKey string, signType string) string {
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
	buf.WriteString(apiKey)

	var (
		dataMd5    [16]byte
		dataSha256 []byte
		str        string
	)

	switch signType {
	case "MD5":
		dataMd5 = md5.Sum(buf.Bytes())
		str = hex.EncodeToString(dataMd5[:]) //需转换成切片
	case "HMAC-SHA256":
		h := hmac.New(sha256.New, []byte(apiKey))
		h.Write(buf.Bytes())
		dataSha256 = h.Sum(nil)
		str = hex.EncodeToString(dataSha256[:])
	}

	return strings.ToUpper(str)
}

// 验证签名
func ValidSign(params map[string]string, apiKey string, signKey string, signType string) bool {
	if _, ok := params[signKey]; !ok {
		return false
	}
	return params[signKey] == Sign(params, apiKey, signType)
}
