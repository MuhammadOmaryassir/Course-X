package entity

type Video struct {
	ID          int64  `json:"id"`
	URL         string `json:"url"`
	Description string `json:"description"`
}
