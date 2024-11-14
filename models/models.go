package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id           int32  `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
	Time   string `json:"time"`
}

type Message struct {
	OwnerId     string `json:"ownerId"`
	ForWhoId    string `json:"forWhoId"`
	TextMessage string `json:"textMessage"`
	TimeSand    string `json:"timeSand"`
	JwtToken    string `json:"jwtToken"`
}

//
//type UserInDB struct {
//	Email    string `json:"email"`
//	PassHash string `json:"password_hash"`
//}

type ResultToUser struct {
	Success bool   `json:"success"`
	Type    string `json:"type"`
}
