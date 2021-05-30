package main

import "fmt"

func main() {
	var flag int

	for true {
		fmt.Println("请输入选择:")
		fmt.Scan(&flag)
		switch flag {
		case 1:
			fmt.Println("11111111111111")
		case 2:
			fmt.Println(222222222222)
		case 3:
			fmt.Println(333333333333)
		default:
			fmt.Println("nnnnnnnnnnnnn")
		}
	}
}
