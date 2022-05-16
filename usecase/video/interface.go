package video

import (
	"github.com/Rodrigueslcs/challenge-backend/entity"
)

type Repository interface {
	List() ([]*entity.Video, error)
	Get(id int) (*entity.Video, error)
	Create(e *entity.Video) (int, error)
	Update(e *entity.Video) error
	Delete(id int) error
}
type UseCase interface {
	ListVideos() ([]*entity.Video, error)
	GetVideo(id int) (*entity.Video, error)
	CreateVideo(title, description, url string) (int, error)
	UpdateVideo(id int, title, description, url string) error
	DeleteVideo(id int) error
}
