package mypackage

import "fmt"

type Student struct{
	Name string
	Age int
}

func New(){
	fmt.Println("mypackage.New")
}
