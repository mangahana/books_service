package dto

type AddBook struct {
	Link          string   `json:"link" validate:"required"`
	Name          string   `json:"name" validate:"required"`
	OriginalTitle string   `json:"original_title" validate:"required"`
	Description   string   `json:"descrpition" validate:"required"`
	Poster        string   `json:"poster" validate:"required,base64"`
	Genres        []string `json:"genres" validate:"required,dive,gte=1"`
	Authors       []string `json:"authors" validate:"required,dive,gte=1"`
	Artists       []string `json:"artists" validate:"required,dive,gte=1"`
	Formats       []string `json:"formats" validate:"required,dive,gte=1"`
	StatusId      string   `json:"status_id" validate:"required,dive,gte=1"`
	TypeId        string   `json:"type_id" validate:"required,dive,gte=1"`
	ReleaseDate   string   `json:"release_date" validate:"required,datetime"`
	OwnerTeamID   int      `json:"owner_team_id"`
}
