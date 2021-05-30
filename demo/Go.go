package main

import (
	"fmt"
	"module1/demo/utils"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
	utils.Test()
}
