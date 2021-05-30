package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"module1/mypackage"
	"net/http"
)

func main() {
	mypackage.New()
	a := mypackage.Student{"huangliang",123}
	fmt.Println(a)
	ginServ := gin.Default()
	ginServ.Any("/aa", WebRoot)
	ginServ.Any("/bb", WebRoot)
	ginServ.Run(":8888")
}

func WebRoot(context *gin.Context) {
	fmt.Println(context.Request)
	context.String(http.StatusOK, "hello, world")
}