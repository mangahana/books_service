package dto

type AddBook struct {
	Link          string `json:"link" validate:"required"`
	Name          string `json:"name" validate:"required"`
	OriginalTitle string `json:"original_title" validate:"required"`
	Description   string `json:"descrpition" validate:"required"`
	Poster        string `json:"poster" validate:"required,base64"`
	Genres        []int  `json:"genres" validate:"required,dive,gte=1"`
	Authors       []int  `json:"authors" validate:"required,dive,gte=1"`
	Artists       []int  `json:"artists" validate:"required,dive,gte=1"`
	StatusId      int    `json:"status_id" validate:"required,dive,gte=1"`
	TypeId        int    `json:"type_id" validate:"required,dive,gte=1"`
	ReleaseDate   string `json:"release_date" validate:"required,datetime"`
	OwnerTeamID   int    `json:"owner_team_id"`
}
