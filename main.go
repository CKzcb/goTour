/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/3/10 11:00
 */

package main

import (
	"github.com/CKzcb/goTour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("cmd.Execute err: %v", err)
	}
}
