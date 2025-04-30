package application

import (
	"server/domain"
)

type UserService interface {
	FindUserByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
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

func (s *UserServiceImpl) CreateUser(user *domain.User) error {
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) UpdateUser(user *domain.User) error {
	err := s.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
