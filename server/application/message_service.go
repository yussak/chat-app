package application

import (
	"database/sql"
	"server/domain"
)

type MessageService interface {
	ListMessagesByChannelID(channelID string) ([]domain.Message, error)
	AddMessage(content string, channelID int, userID int) (domain.Message, error)
	DeleteMessageAndRelationData(id string, tx *sql.Tx) error
}

type MessageServiceImpl struct {
	// todo:messageRepoにするかも
	repo         domain.MessageRepository
	reactionRepo domain.ReactionRepository
}

func NewMessageService(repo domain.MessageRepository, reactionRepo domain.ReactionRepository) MessageService {
	return &MessageServiceImpl{repo: repo, reactionRepo: reactionRepo}
}

func (s *MessageServiceImpl) ListMessagesByChannelID(channelID string) ([]domain.Message, error) {
	messages, err := s.repo.FindByChannelID(channelID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *MessageServiceImpl) AddMessage(content string, channelID int, userID int) (domain.Message, error) {
	message, err := s.repo.AddMessage(content, channelID, userID)
	if err != nil {
		return domain.Message{}, err
	}
	return message, nil
}

func (s *MessageServiceImpl) DeleteMessageAndRelationData(id string, tx *sql.Tx) error {
	err := s.reactionRepo.Delete(id, tx)
	if err != nil {
		return err
	}

	err = s.repo.Delete(id, tx)
	if err != nil {
		return err
	}

	return nil
}
