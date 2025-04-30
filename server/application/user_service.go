package application

import (
	"database/sql"
	"server/domain"
)

type UserService interface {
	FindUserByEmail(email string) (*domain.User, error)
	CreateUser(db *sql.DB, user *domain.User) error
	UpdateUser(db *sql.DB, user *domain.User) error
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

func (s *UserServiceImpl) CreateUser(db *sql.DB, user *domain.User) error {
	err := s.userRepo.CreateUser(db, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) UpdateUser(db *sql.DB, user *domain.User) error {
	err := s.userRepo.UpdateUser(db, user)
	if err != nil {
		return err
	}

	return nil
}
