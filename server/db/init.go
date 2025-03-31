package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

func Init() {
	// データベース接続設定
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=Asia/Tokyo",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

	err = initTable()
	if err != nil {
		log.Fatalf("テーブル初期化エラー: %v", err)
	}
}

func initTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE NOT NULL,
		image TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS workspaces (
		id SERIAL PRIMARY KEY,
		owner_id INTEGER REFERENCES users(id),
		name TEXT NOT NULL,
		theme TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS workspace_members (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id),
		workspace_id INTEGER REFERENCES workspaces(id),
		display_name TEXT NOT NULL,
		image_url TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS channels (
		id SERIAL PRIMARY KEY,
		workspace_id INTEGER REFERENCES workspaces(id),
		name TEXT NOT NULL,
		is_public BOOLEAN NOT NULL DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS channel_members (
		id SERIAL PRIMARY KEY,
		channel_id INTEGER REFERENCES channels(id),
		user_id INTEGER REFERENCES users(id),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		channel_id INTEGER REFERENCES channels(id),
		content TEXT NOT NULL,
		user_id INTEGER REFERENCES users(id),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS reactions (
		id SERIAL PRIMARY KEY,
		message_id INTEGER REFERENCES messages(id),
		user_id INTEGER REFERENCES users(id),
		emoji TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (message_id, user_id, emoji)
	);

  `

	_, err := DB.Exec(query)
	log.Println("テーブルのセットアップ完了")

	return err
}
