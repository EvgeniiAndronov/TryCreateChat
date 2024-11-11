package models

type User struct {
	Id           int32
	Email        string
	PasswordHash string
	LastUpdate   int64
}

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
	Time   string `json:"time"`
}
