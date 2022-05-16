package video

import "github.com/Rodrigueslcs/challenge-backend/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ListVideos() ([]*entity.Video, error) {
	return s.repo.List()
}

func (s *Service) GetVideo(id int) (*entity.Video, error) {
	return s.repo.Get(id)
}

func (s *Service) CreateVideo(title, description, url string) (int, error) {
	e, err := entity.NewVideo(title, description, url)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

func (s *Service) UpdateVideo(id int, title, description, url string) error {
	e, err := entity.NewVideo(title, description, url)
	if err != nil {
		return err
	}
	return s.repo.Update(e)
}

func (s *Service) DeleteVideo(id int) error {
	_, err := s.GetVideo(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
