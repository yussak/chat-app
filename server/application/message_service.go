package application

import (
	"database/sql"
	"server/domain"
)

type MessageService interface {
	ListMessages(channelID string) ([]domain.Message, error)
	AddMessage(content string, channelID int, userID int) (domain.Message, error)
	DeleteMessageAndRelationData(id string, currentUserID string, tx *sql.Tx) error
}

type MessageServiceImpl struct {
	messageRepo  domain.MessageRepository
	reactionRepo domain.ReactionRepository
}

func NewMessageService(messageRepo domain.MessageRepository, reactionRepo domain.ReactionRepository) MessageService {
	return &MessageServiceImpl{messageRepo: messageRepo, reactionRepo: reactionRepo}
}

func (s *MessageServiceImpl) ListMessages(channelID string) ([]domain.Message, error) {
	messages, err := s.messageRepo.FindByChannelID(channelID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *MessageServiceImpl) AddMessage(content string, channelID int, userID int) (domain.Message, error) {
	message, err := s.messageRepo.AddMessage(content, channelID, userID)
	if err != nil {
		return domain.Message{}, err
	}
	return message, nil
}

func (s *MessageServiceImpl) DeleteMessageAndRelationData(id string, currentUserID string, tx *sql.Tx) error {
	err := s.reactionRepo.Delete(id, tx)
	if err != nil {
		return err
	}

	err = s.messageRepo.Delete(id, currentUserID, tx)
	if err != nil {
		return err
	}

	return nil
}
