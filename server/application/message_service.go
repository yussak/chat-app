package application

import (
	"server/domain"
)

type MessageService interface {
	ListMessagesByChannelID(channelID string) ([]domain.Message, error)
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
