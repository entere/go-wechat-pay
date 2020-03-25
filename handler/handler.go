/**
 * @Author: entere
 * @Description:
 * @File:  handler
 * @Version: 1.0.0
 * @Date: 2020/3/25 22:04
 */

package handler

import "github.com/entere/go-wechat-pay/mch/core"

func GetApiKey() string {
	if core.IsSandbox {
		return core.WX_API_KEY_SANBOX
	}
	return core.WX_API_KEY
}
