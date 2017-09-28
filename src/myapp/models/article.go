package models

type Article struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Category        string `json:"category"`
	Banner_url      string `json:"banner_url"`
	Content         string `json:"content"`
	Created         string `json:"created"`
	Author_position string `json:"author_position"`
	Author_id       int    `json:"author_id"`
}
