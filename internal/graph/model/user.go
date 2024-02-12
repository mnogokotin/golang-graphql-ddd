package model

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	PostIDs []*string `json:"postIds"`
}
