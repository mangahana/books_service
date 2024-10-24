package models

type User struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	Photo       string   `json:"photo"`
	IsBanned    bool     `json:"is_banned"`
	Permissions []string `json:"permissions"`
}
