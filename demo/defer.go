package main

import (
	"fmt"
)

func main(){//defer会在函数执行完毕之后执行
	defer fmt.Println("finish")
	defer fmt.Println("finish1")
	defer fmt.Println("finish2")
	defer fmt.Println("finish3")
	fmt.Println("start")
}
/**
	start
	finish3
	finish2
	finish1
	finish
 */