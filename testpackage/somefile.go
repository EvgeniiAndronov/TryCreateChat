package testpackage

import "github.com/gin-gonic/gin"

type Obj struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var testObj = []Obj{
	Obj{Id: 1, Name: "Jhon"},
	Obj{Id: 2, Name: "Smith"},
}

func SomeInt(c *gin.Context) {
	c.JSON(200, testObj)
}
