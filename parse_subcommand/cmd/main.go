/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/3/10 11:19
 */

package main

import (
	"flag"
	"fmt"
	"github.com/CKzcb/goTour/parse_subcommand"
)

/**
* t1
*  @Description:
* 	flag.Parse()
*		1. 所有命令行参数注册的最后调用
*		2. 解析并绑定命令行参数
 */
func t1() {

	var name string

	//flag.StringVar(&name, "name", "", " please input name")
	// 1. to parse, get args
	flag.Parse()
	// 2. create  subcommand cmd
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "", "")

	pyCmd := flag.NewFlagSet("py", flag.ExitOnError)
	pyCmd.StringVar(&name, "n", "", "")

	// 3. need to parse
	args := flag.Args()
	fmt.Println("args : ", args)
	if len(args) <= 0 {
		return
	}
	// 4. select son cmd
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "py":
		_ = pyCmd.Parse(args[1:])
	}
	fmt.Println("input name is : ", name)
}

func t2() {
	var name parse_subcommand.Name
	flag.Var(&name, "name", "name")
	flag.Parse()
	fmt.Println("in t2 ... ", name)
}

func main() {
	//t1()
	t2()
}
