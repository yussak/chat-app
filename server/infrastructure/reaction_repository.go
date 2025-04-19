package infrastructure

import (
	"database/sql"
)

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
