package models

import (
	"database/sql"
	"log"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Image     string `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FindUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `SELECT id, name, email, image, created_at, updated_at FROM users WHERE email = $1`
	user := &User{}

	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Image, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil // 初回ログインの場合はnil
	} else if err != nil {
		log.Printf("ユーザー検索失敗: %v", err)
		return nil, err
	}

	return user, nil
}

func CreateUser(db *sql.DB, user *User) error {
	query := `INSERT INTO users (name, email, image) VALUES ($1, $2, $3) RETURNING id`
	return db.QueryRow(query, user.Name, user.Email, user.Image).Scan(&user.ID)
}

func UpdateUser(db *sql.DB, user *User) error {
	query := `UPDATE users SET name=$1, image=$2, updated_at=NOW() WHERE email=$3 RETURNING id`
	return db.QueryRow(query, user.Name, user.Image, user.Email).Scan(&user.ID)
}