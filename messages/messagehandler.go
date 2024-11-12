package messages

import (
	"ChatTryCreate/models"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var Messages []models.Message

func MessageRecive(c *gin.Context) {
	message := models.Message{}
	db, err := gorm.Open(sqlite.Open("chat_try.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	message.OwnerId = c.Query("OwnerId")
	message.ForWhoId = c.Query("ForWhoId")
	message.TextMessage = c.Query("TextMessage")
	message.TimeSand = c.Query("TimeSand")
	message.JwtToken = c.Query("JwtToken")

	//fmt.Println(message)

	if checkToken(message) {
		err = addMsgToDb(message, db)
		if err != nil {
			log.Fatal(err)
		}

		err = sandToUserMsg(message)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Wrong token")
	}

}

func algTokenDe(token string, ownerId string, time string) bool {
	must_be_token := (len(ownerId)) * (len(ownerId))
	str_mbt := strconv.Itoa(must_be_token)

	buffer := bytes.Buffer{}
	buffer.WriteString(str_mbt)
	buffer.WriteString(time)

	new_mbt := buffer.String()

	fmt.Println(token, new_mbt)

	if new_mbt == token {
		return true
	} else {
		return false
	}
}

func checkToken(message models.Message) bool {
	return algTokenDe(message.JwtToken, message.OwnerId, message.TimeSand)
}

func addMsgToDb(message models.Message, db *gorm.DB) error {
	Messages = append(Messages, message)
	db.Create(&message)
	return nil
}

func sandToUserMsg(message models.Message) error {

	return nil
}

func ShowMessages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Messages": Messages})
}
