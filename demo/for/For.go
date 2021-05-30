package main

import (
	"fmt"
)

func main() {
	//for true {
	//	time.Sleep(1000000)
	//	fmt.Println(1);
	//}

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Print(sum, "->") //0->1->3->6->10->15->21->28->36->45->55->55
	}
	fmt.Println(sum)

	numberArr := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(numberArr) //[1 2 3 4 5 6]

	for i := 0; i < len(numberArr); i++ {
		fmt.Print(numberArr[i], ",") //1,2,3,4,5,6,
	}

	for i, x := range numberArr {
		fmt.Printf("第%d位值为：%d\n", i, x) //Printf=> format(String,a,b,c)
	}
}
