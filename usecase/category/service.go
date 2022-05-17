package category

import "github.com/Rodrigueslcs/challenge-backend/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ListCategories() ([]*entity.Category, error) {
	return s.repo.List()
}

func (s *Service) GetCategory(id int) (*entity.Category, error) {
	return s.repo.Get(id)
}

func (s *Service) CreateCategory(title, color string) (int, error) {

	e, err := entity.NewCategory(title, color)
	if err != nil {
		return e.ID, err
	}
	err = e.Validate()
	if err != nil {
		return 0, err
	}
	return s.repo.Create(e)
}

func (s *Service) UpdateCategory(id int, title, color string) error {
	e, err := entity.NewCategory(title, color)
	if err != nil {
		return err
	}
	err = e.Validate()
	if err != nil {
		return err
	}
	return s.repo.Update(e)
}

func (s *Service) DeleteCategory(id int) error {
	_, err := s.GetCategory(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
