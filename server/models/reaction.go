package models

import "time"

type Reaction struct {
	ID int
	MessageID int
	UserID int
	Emoji string
	CreatedAt time.Time
}