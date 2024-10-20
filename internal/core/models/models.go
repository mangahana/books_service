package models

import "time"

type BookType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID     int    `json:"id"`
	Link   string `json:"link"`
	Name   string `json:"name"`
	Poster string `json:"poster"`
	Type   string `json:"type"`
}

type OneBook struct {
	ID            int       `json:"id"`
	Link          string    `json:"link"`
	Name          string    `json:"name"`
	OriginalTitle string    `json:"original_title"`
	Description   string    `json:"description"`
	Poster        string    `json:"poster"`
	Type          string    `json:"type"`
	Genres        []string  `json:"genres"`
	Authors       []string  `json:"authors"`
	Artists       []string  `json:"artists"`
	Status        string    `json:"status"`
	ReleaseDate   time.Time `json:"release_date"`
}

type Chapter struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Volume    int       `json:"volume"`
	Number    string    `json:"number"`
	CreatedAt time.Time `json:"created_at"`
}
