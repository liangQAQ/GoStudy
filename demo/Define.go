package main

import "fmt"

var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明

type param map[string]string

func test() {
	var available bool            // 一般声明
	valid := false                // 简短声明
	available = true              // 赋值操作
	numbers := [6]int{1, 2, 3, 5} //数组
	fmt.Println(available, valid, numbers)
	//param["112233"] = "223344"
	//test()
	//a1(param)
}

func a1(a map[string]string) {
	fmt.Println(a)
}

func main() {
	test()
	//a();
}
