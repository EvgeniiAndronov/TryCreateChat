package regauth

import (
	"ChatTryCreate/models"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func MD5(pas string) string {
	h := md5.Sum([]byte(pas))
	return fmt.Sprintf("%x", h)
}

func RegistrationUser(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	hash := MD5(password)

	db, err := gorm.Open(sqlite.Open("chat_try.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database (add user connection)")
	}

	user := models.User{Email: email, PasswordHash: hash}
	//var exist bool
	count := db.Where("email = ?", email).Find(&user).RowsAffected
	if count > 0 {
		fmt.Println("уже существовал")
		fmt.Println(count)
		c.JSON(http.StatusOK, models.ResultToUser{Success: false, Type: "User has been registered"})
	} else {
		fmt.Println("создан---")
		fmt.Println(count)
		db.Create(&user)
		c.JSON(http.StatusOK, models.ResultToUser{Success: true, Type: "Created User"})
	}
}

func SuccesRegistration(c *gin.Context) {
	c.JSON(http.StatusOK, models.ResultToUser{Success: true, Type: "Created User"})
}
