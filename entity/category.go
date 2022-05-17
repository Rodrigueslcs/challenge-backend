package entity

import "errors"

type Category struct {
	ID    int
	Title string
	Color string
}

func NewCategory(title, color string) (*Category, error) {
	c := &Category{
		Title: title,
		Color: color,
	}

	return c, nil
}

func (c *Category) Validate() error {
	if len(c.Title) < 1 {
		return errors.New("TODOS OS CAMPOS SÃO OBRIGATORIO")
	}
	return nil
}
