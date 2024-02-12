package model

type Post struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"userId"`
	User   *User  `json:"user"`
}
