package application

import (
	"server/domain"
)

type ChannelService interface {
	GetChannel(id string) (domain.Channel, error)
}

type ChannelServiceImpl struct {
	channelRepo domain.ChannelRepository
}

func NewChannelService(channelRepo domain.ChannelRepository) ChannelService {
	return &ChannelServiceImpl{channelRepo: channelRepo}
}

func (s *ChannelServiceImpl) GetChannel(id string) (domain.Channel, error) {
	channel, err := s.channelRepo.GetChannel(id)
	if err != nil {
		return domain.Channel{}, err
	}

	return channel, nil
}
