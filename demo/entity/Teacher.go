package entity

type Teacher struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}
