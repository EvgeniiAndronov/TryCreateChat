package main

import (
	"ChatTryCreate/messages"
	"ChatTryCreate/models"
	"ChatTryCreate/posts"
	"ChatTryCreate/regauth"
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

	db.AutoMigrate(&models.User{}, &models.Message{})

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/posts", posts.ShowPostsOnPage)
	router.GET("/postjson", posts.ShowPostsJson)
	router.POST("/message", messages.MessageRecive)
	router.GET("/showMessage", messages.ShowMessages)
	router.POST("/reg", regauth.RegistrationUser)
	router.GET("Reg_succ", regauth.SuccesRegistration)

	router.Run(":8080")
}
