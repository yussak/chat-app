package application

import (
	"server/domain"
)

type ReactionService interface {
	ListReactions(messageId string) ([]domain.Reaction, error)
	AddReaction(messageId string, userId int, emoji string) error
}

type ReactionServiceImpl struct {
	reactionRepo domain.ReactionRepository
}

func NewReactionService(reactionRepo domain.ReactionRepository) ReactionService {
	return &ReactionServiceImpl{reactionRepo: reactionRepo}
}

func (s *ReactionServiceImpl) ListReactions(messageId string) ([]domain.Reaction, error) {
	reactions, err := s.reactionRepo.ListReactions(messageId)
	if err != nil {
		return nil, err
	}

	return reactions, nil
}

func (s *ReactionServiceImpl) AddReaction(messageId string, userId int, emoji string) error {
	err := s.reactionRepo.AddReaction(messageId, userId, emoji)
	if err != nil {
		return err
	}

	return nil
}
