package domain

type Post struct {
	ID   string `json:"id"`
	User User   `json:"user"`
	Text string `json:"text"`
}
