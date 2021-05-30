package main

import "fmt"

/**
func function_name( [parameter list] ) (return_types) {
   函数体
}
*/
func main() {
	a, b := 1, 3
	//swap(a,b)//改变不了
	a, b = swap1(a, b)
	fmt.Println(a, b)
}

func swap(a int, b int) (int, int) { //基本类型 交换不了
	var temp = a
	a = b
	b = temp
	return a, b
}

func swap1(a int, b int) (int, int) {
	return b, a
}
