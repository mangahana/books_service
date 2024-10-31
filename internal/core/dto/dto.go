package dto

type AddBook struct {
	Link          string   `json:"link" validate:"required"`
	Name          string   `json:"name" validate:"required"`
	OriginalTitle string   `json:"original_title" validate:"required"`
	Description   string   `json:"descrpition" validate:"required"`
	Poster        string   `json:"poster" validate:"required,base64"`
	Genres        []string `json:"genres" validate:"required,dive"`
	Authors       []string `json:"authors" validate:"required,dive"`
	Artists       []string `json:"artists" validate:"required,dive"`
	Formats       []string `json:"formats" validate:"required,dive"`
	StatusId      string   `json:"status_id" validate:"required,dive"`
	TypeId        string   `json:"type_id" validate:"required,dive"`
	ReleaseDate   string   `json:"release_date" validate:"required,datetime"`
	OwnerTeamID   int      `json:"owner_team_id"`
}

type CreateDraft struct {
	BookID int `json:"book_id" validate:"required,number,gte=1"`
	TeamId int `json:"team_id" validate:"required,number,gte=1"`
}

type UploadPage struct {
	ChapterID string `json:"chapter_id" validate:"required"`
	Image     string `json:"image" validate:"required,base64"`
}

type CreateChapter struct {
	ID     string   `json:"id"`
	BookID int      `json:"book_id"`
	Volume int      `json:"volume"`
	Number string   `json:"number"`
	Name   string   `json:"name"`
	Pages  []string `json:"pages"`
}

type AddPerson struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
