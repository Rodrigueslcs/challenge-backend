package category

import (
	"github.com/Rodrigueslcs/challenge-backend/entity"
)

type Repository interface {
	List() ([]*entity.Category, error)
	Get(id int) (*entity.Category, error)
	Create(e *entity.Category) (int, error)
	Update(e *entity.Category) error
	Delete(id int) error
}
type UseCase interface {
	ListCategories() ([]*entity.Category, error)
	GetCategory(id int) (*entity.Category, error)
	CreateCategory(title, color string) (int, error)
	UpdateCategory(id int, title, color string) error
	DeleteCategory(id int) error
}
