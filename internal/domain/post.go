package domain

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"user"`
	Text   string `json:"text"`
}
