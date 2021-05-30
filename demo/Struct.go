package main

import "fmt"

type Student struct {
	name   string
	age    int
	number int
}

func main() {
	var a = Student{"黄亮", 2, 20120304}
	var b = Student{name: "黄亮", number: 20120507}

	var c = [2]Student{}
	c[0] = a
	c[1] = b

	fmt.Println(c)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(b.number)

	print(b)
}

func print(s Student) {
	fmt.Println(s)
}
