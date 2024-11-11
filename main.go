package main

import (
	"ChatTryCreate/models"
	"ChatTryCreate/posts"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var router *gin.Engine

func main() {

	db, err := gorm.Open(sqlite.Open("chat_try.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/posts", posts.ShowPostsOnPage)
	router.GET("/postjson", posts.ShowPostsJson)

	router.Run(":8080")

}
