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

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	)`
	_, err := DB.Exec(query)
	log.Println("テーブルのセットアップ完了")

	return err
}
