package main

import "fmt"

func main() {
	var name string
	var age int
	var salary float32
	//带"&"为地址，传地址才会影响外面的值
	fmt.Scanln(&name) //程序执行到这会阻塞 等待用户输入
	fmt.Scanln(&age)
	fmt.Scanln(&salary)
	fmt.Print(age, name, salary)
}
