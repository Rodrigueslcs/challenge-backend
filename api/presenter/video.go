package presenter

type Video struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	CategoryID  int    `json:"category_id"`
}
