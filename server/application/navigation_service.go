package application

import "server/domain"

type NavigationService interface {
	ListSidebarProps() ([]domain.NavigationSidebarProps, error)
}

type navigationServiceImpl struct {
	repo domain.NavigationRepository
}

func NewNavigationService(repo domain.NavigationRepository) NavigationService {
	return &navigationServiceImpl{repo: repo}
}

func (s *navigationServiceImpl) ListSidebarProps() ([]domain.NavigationSidebarProps, error) {
	props, err := s.repo.GetSidebarProps()
	if err != nil {
		return nil, err
	}
	return props, nil
}