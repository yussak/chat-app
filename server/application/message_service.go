package application

import (
	"server/domain"
)

type MessageService interface {
	GetMessages(channelID string) ([]domain.Message, error)
}

type MessageServiceImpl struct {
	repo domain.MessageRepository
}

func NewMessageService(repo domain.MessageRepository) MessageService {
	return &MessageServiceImpl{repo: repo}
}

func (s *MessageServiceImpl) GetMessages(channelID string) ([]domain.Message, error) {
	messages, err := s.repo.Get(channelID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
