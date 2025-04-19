package infrastructure

import (
	"database/sql"
)

// todo:ほかもこんな感じにレシーバなど改善
type ReactionRepository struct{}

func NewReactionRepository() *ReactionRepository {
	return &ReactionRepository{}
}

func (r *ReactionRepository) Delete(id string, tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM reactions WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
