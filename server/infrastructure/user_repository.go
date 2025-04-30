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
