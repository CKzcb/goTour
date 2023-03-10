/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: test_main
 * @Version: 1.0.0
 * @Date: 2023/3/10 15:39
 */

package main

import (
	"fmt"
	"reflect"
)

type T1User struct {
	Name string
	Age  int
}

func TestType1(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("a in string ... ")
	case int:
		fmt.Println("a is int ... ")
	case bool:
		fmt.Println("a is bool ")
	case T1User:
		fmt.Println("a is t1 user ... ")
	default:
		fmt.Println("unknown type ... ")
	}
}

type Int int

func (i Int) S1() {
	fmt.Println("i execute s1. ..")
}

func main() {
	var a Int = 20
	TestType1(a)

	c := reflect.TypeOf(a)
	fmt.Println(c.Method(0).Type)
	c.Method(0).Func.Call([]reflect.Value{reflect.ValueOf(a)})

	d := reflect.ValueOf(a)
	d.Method(0).Call([]reflect.Value{})
}
