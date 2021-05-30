package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"module1/demo/entity"
	"net/http"
)

func main() {
	r := gin.Default()

	// 匹配 /user/geektutu
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 匹配users?name=xxx&role=xxx，role可选
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	// POST
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// GET 和 POST 混合
	r.GET("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	//对象绑定
	r.Any("/teacher", func(context *gin.Context) {
		var t entity.Teacher
		//err := context.ShouldBindQuery(&t)
		err := context.ShouldBindUri(&t)
		if err != nil {
			fmt.Println("teacher:", t)
		}
	})

	r.Run(":8888")
}
