package user

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, user *User) error {
	return s.repo.Create(user)
}

func (s *Service) GetUser(ctx context.Context, id uint) (*User, error) {
	return s.repo.GetbyID(id)
}

func (s *Service) UpdateUser(ctx context.Context, user *User) error {
	return s.repo.Update(user)
}

func (s *Service) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.Delete(id)
}

func (s *Service) ListUser(ctx context.Context) ([]User, error) {
	return s.repo.List()
}
