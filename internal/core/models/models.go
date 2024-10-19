package models

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
