package application

import (
	"server/domain"
	"server/models"
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

// todo:一時的にmodelsにしてるので適切にする
// todo:そもそもUser丸ごと渡す必要ないかもしれないので確認
func (s *MessageServiceImpl) AddMessage(content string, channelID int, user models.User) (domain.Message, error ){
  message, err := s.repo.AddMessage(content, channelID, user)
  if err != nil {
    return domain.Message{}, err
  }
  return message, nil
}