package models

type User struct {
	ID          int      `json:"user_id"`
	IsBanned    bool     `json:"is_banned"`
	Permissions []string `json:"permissions"`
}
