/**
 * @Author: entere
 * @Description:
 * @File:  nonce_str
 * @Version: 1.0.0
 * @Date: 2020/3/24 13:19
 */

package wxutils

import (
	"fmt"
	"math/rand"
	"time"
)

//生成8位随机数字
func NonceStr(n int) string {
	if n == 0 {
		n = 8
	}
	leterset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var leteridxbits uint = 6
	var mask int64 = 1<<leteridxbits - 1

	rnd := rand.NewSource(time.Now().UnixNano())
	res := make([]byte, 0, n)

	for i, bits := 0, rnd.Int63(); i < n; i++ {
		if bits == 0 {
			bits = rnd.Int63()
		}
		idx := int(bits & mask)
		if idx < len(leterset) {
			res = append(res, leterset[idx])
		} else {
			i--
		}
		bits >>= leteridxbits
	}
	return string(res[:n])

	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
