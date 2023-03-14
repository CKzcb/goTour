/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: timer_test_main
 * @Version: 1.0.0
 * @Date: 2023/3/10 16:21
 */

package main

import (
	"fmt"
	"github.com/CKzcb/goTour/internal/timer"
)

func main() {
	now := timer.GetNowTime()
	fmt.Println(now)
	fmt.Println(timer.GetCalculateTime(now, "2m"))
}
