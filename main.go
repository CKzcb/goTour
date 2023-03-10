/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/3/10 11:00
 */

package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string

	//flag.StringVar(&name, "name", "", " please input name")
	flag.Parse()

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "", "")

	pyCmd := flag.NewFlagSet("py", flag.ExitOnError)
	pyCmd.StringVar(&name, "n", "", "")

	fmt.Println("input name is : ", name)
}
