package application

import (
	"server/domain"
)

type ReactionService interface {
	ListReactions(messageId string) ([]domain.Reaction, error)
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
