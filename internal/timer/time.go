/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: time
 * @Version: 1.0.0
 * @Date: 2023/3/10 16:17
 */

package timer

import (
	"time"
)

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(CurrentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return CurrentTime.Add(duration), nil
}
