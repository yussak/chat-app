package models

import "time"

type Message struct {
	ID        int
	Content   string
	User      User
	CreatedAt time.Time
	Reactions string `json:"reactions"`
}
