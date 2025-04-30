package infrastructure

import (
	"database/sql"
	"log"
	"server/db"
	"server/domain"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	query := `SELECT id, name, email, image, created_at, updated_at FROM users WHERE email = $1`
	user := &domain.User{}

	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Image, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil // 初回ログインの場合はnil
	} else if err != nil {
		log.Printf("ユーザー検索失敗: %v", err)
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (name, email, image) VALUES ($1, $2, $3) RETURNING id`
	return db.DB.QueryRow(query, user.Name, user.Email, user.Image).Scan(&user.ID)
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	query := `UPDATE users SET name=$1, image=$2, updated_at=NOW() WHERE email=$3 RETURNING id`
	return db.DB.QueryRow(query, user.Name, user.Image, user.Email).Scan(&user.ID)
}
