package infrastructure

import (
	"server/db"
	"server/domain"
)

type ChannelRepository struct{}

func NewChannelRepository() *ChannelRepository {
	return &ChannelRepository{}
}

func (r *ChannelRepository) GetChannel(id string) (domain.Channel, error) {
	query := `SELECT id, name, created_at, updated_at FROM channels WHERE id = $1`

	var channel domain.Channel
	err := db.DB.QueryRow(query, id).Scan(
		&channel.ID,
		&channel.Name,
		&channel.CreatedAt,
		&channel.UpdatedAt,
	)
	if err != nil {
		return domain.Channel{}, err
	}

	return channel, nil
}
