package application

import (
	"server/domain"
)

type MessageService interface {
	ListMessagesByChannelID(channelID string) ([]domain.Message, error)
	AddMessage(content string, channelID int, userID int) (domain.Message, error)
}

type MessageServiceImpl struct {
	repo domain.MessageRepository
}

func NewMessageService(repo domain.MessageRepository) MessageService {
	return &MessageServiceImpl{repo: repo}
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
