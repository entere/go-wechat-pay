/**
 * @Author: entere
 * @Description:
 * @File:  time_test
 * @Version: 1.0.0
 * @Date: 2020/3/27 20:20
 */

package wxutils

import (
    "testing"
    "time"
)

func TestTimeToString(t *testing.T) {
    toString := TimeToString(time.Now())
    t.Logf(toString)
}