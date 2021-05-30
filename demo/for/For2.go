package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a int
	for true {
		rand.Seed(time.Now().Unix()) //设置随机种子
		a = rand.Intn(100)
		fmt.Println(a)
		if a > 80 {
			break
		}
	}
}
