package model

type Post struct {
	Title   string   `json:"title"`
	Section []string `json:"sections"`
	Date    string   `json:"date"`
	page    *Page
}
