package model

type Book struct {
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Year   int    `json:"year,omitempty"`
}
