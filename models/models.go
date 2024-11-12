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

type Message struct {
	OwnerId     string `json:"ownerId"`
	ForWhoId    string `json:"forWhoId"`
	TextMessage string `json:"textMessage"`
	TimeSand    string `json:"timeSand"`
	JwtToken    string `json:"jwtToken"`
}
