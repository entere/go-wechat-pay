/**
 * @Author: entere
 * @Description:
 * @File:  time
 * @Version: 1.0.0
 * @Date: 2020/3/19 15:19
 */

package wxutils

import "time"

// FormatTime 将参数 t 格式化成北京时间 yyyyMMddHHmmss.
func TimeToString(t time.Time) string {
	beijingLocation := time.FixedZone("Asia/Shanghai", 8*60*60)
	return t.In(beijingLocation).Format("20060102150405")
}

// ParseTime 将北京时间 yyyyMMddHHmmss 字符串解析到 time.Time.
func StringToTime(value string) (time.Time, error) {
	beijingLocation := time.FixedZone("Asia/Shanghai", 8*60*60)
	return time.ParseInLocation("20060102150405", value, beijingLocation)
}
