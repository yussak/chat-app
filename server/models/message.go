package models

type Message struct {
	ID   int
	Content string
	User User
	Reactions string `json:"reactions"`
}
