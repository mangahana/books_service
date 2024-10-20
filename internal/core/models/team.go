package models

type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Photo   string `json:"photo"`
	Type    string `json:"type"`
	OwnerId int    `json:"owner_id"`
}

type TeamMember struct {
	UserID      int      `json:"user_id"`
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
}
