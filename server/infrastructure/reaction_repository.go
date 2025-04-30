package infrastructure

import (
	"database/sql"
	"server/db"
	"server/domain"
)

type ReactionRepository struct{}

func NewReactionRepository() *ReactionRepository {
	return &ReactionRepository{}
}

func (r *ReactionRepository) Delete(id string, tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM reactions WHERE message_id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ReactionRepository) ListReactions(messageId string) ([]domain.Reaction, error) {
	rows, err := db.DB.Query(`
	SELECT emoji, COUNT(*)
	FROM reactions
	WHERE message_id = $1
	GROUP BY emoji
	`, messageId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	reactions := []domain.Reaction{}
	for rows.Next() {
		var reaction domain.Reaction
		if err := rows.Scan(&reaction.Emoji, &reaction.Count); err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}

	return reactions, nil
}

func (r *ReactionRepository) AddReaction(messageId string, userId int, emoji string) error {
	// 既存のリアクションを確認
	var exists bool
	err := db.DB.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM reactions 
			WHERE message_id = $1 AND user_id = $2 AND emoji = $3
		)`,
		messageId, userId, emoji,
	).Scan(&exists)

	if err != nil {
		return err
	}

	if exists {
		// リアクションが存在する場合は削除
		_, err = db.DB.Exec(`
	DELETE FROM reactions 
	WHERE message_id = $1 AND user_id = $2 AND emoji = $3`,
			messageId, userId, emoji,
		)
		return err

	}
	// リアクションが存在しない場合は追加
	_, err = db.DB.Exec(`
	INSERT INTO reactions (message_id, user_id, emoji)
		VALUES ($1, $2, $3)`,
		messageId, userId, emoji,
	)

	return err
}
