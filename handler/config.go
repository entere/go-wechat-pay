/**
 * @Author: entere
 * @Description:
 * @File:  handler
 * @Version: 1.0.0
 * @Date: 2020/3/25 22:04
 */

package handler

var isSandbox = false

var (
	appID     string
	appSecret string
	mchID     string
	apiKey    string
	totalFree int64
	notifyUrl string
)


func Init() {

	//todo 不用沙箱  换成配置文件
	if !isSandbox {
		appID     = "x"
		appSecret = "x"
		mchID     = "x"
		apiKey    = "x"
		totalFree = 1
		notifyUrl = "x"
	} else {
		appID     = "x"
		appSecret = "x"
		mchID     = "x"
		apiKey    = "x"
		totalFree = 301
		notifyUrl = "x"
	}


}

