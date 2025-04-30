package domain

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	FindUserByEmail(email string) (*User, error)
	CreateUser(db *sql.DB, user *User) error
	UpdateUser(db *sql.DB, user *User) error
}
