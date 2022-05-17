package entity

import "errors"

type Video struct {
	ID          int
	Title       string
	Description string
	URL         string
	CategoryID  int
}

func NewVideo(title, description, url string, categoryID int) (*Video, error) {
	v := &Video{
		Title:       title,
		Description: description,
		URL:         url,
		CategoryID:  categoryID,
	}

	return v, nil
}

func (v *Video) Validate() error {
	if len(v.Title) < 1 || len(v.Description) < 1 || len(v.URL) < 1 {
		return errors.New("TODOS OS CAMPOS SÃO OBRIGATORIO")
	}
	if v.CategoryID < 1 {
		v.CategoryID = 1
	}
	return nil
}
