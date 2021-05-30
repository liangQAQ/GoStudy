package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("变量的地址: %x\n", &a) //c0000b0058(地址每一次都不一样)
	//var p *int;
	var p *int
	var pp **int
	p = &a
	pp = &p
	fmt.Println("a:", &a)      //0xc0000ac058
	fmt.Println("p:", p)       //0xc0000ac058
	fmt.Println("*p:", *p)     //10
	fmt.Println("**pp:", **pp) //10
	a = 20
	fmt.Println("a:", &a)      //0xc0000ac058
	fmt.Println("p:", p)       //0xc0000ac058
	fmt.Println("*p:", *p)     //20
	fmt.Println("**pp:", **pp) //20
}
