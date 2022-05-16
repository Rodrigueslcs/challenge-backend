package entity

type Video struct {
	ID          int
	Title       string
	Description string
	URL         string
}

func NewVideo(title, description, url string) (*Video, error) {
	v := &Video{
		Title:       title,
		Description: description,
		URL:         url,
	}
	return v, nil
}
