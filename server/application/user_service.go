package application

import (
	"server/domain"
)

type UserService interface {
	FindUserByEmail(email string) (*domain.User, error)
}

type UserServiceImpl struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (s *UserServiceImpl) FindUserByEmail(email string) (*domain.User, error) {
	user, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
