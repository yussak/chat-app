package test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

const (
	testDBHost     = "localhost"
	testDBPort     = "5433"
	testDBUser     = "test_user"
	testDBPassword = "test_password"
	testDBName     = "test_db"
)

func GetTestDB(t *testing.T) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		testDBHost, testDBPort, testDBUser, testDBPassword, testDBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("データベース接続エラー: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("データベース接続テストエラー: %v", err)
	}

	if err := initTestTable(db); err != nil {
		t.Fatalf("テーブル初期化エラー: %v", err)
	}

	return db
}

func initTestTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE NOT NULL,
		image TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
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
	`

	_, err := db.Exec(query)
	return err
}

// TruncateTables は指定されたテーブルのデータを削除します
func TruncateTables(t *testing.T, db *sql.DB, tables ...string) {
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
		if err != nil {
			t.Fatalf("テーブル %s のデータ削除エラー: %v", table, err)
		}
	}
}

// GetTestDBURL はテストデータベースの接続URLを返します
func GetTestDBURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		testDBUser, testDBPassword, testDBHost, testDBPort, testDBName,
	)
}
