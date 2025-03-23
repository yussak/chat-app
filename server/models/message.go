package models

import "time"

type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	User      User   `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Reactions string `json:"reactions"`
}
