package model

import "time"

type MarkdownPost struct {
	Id        int64     `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ShowMarkdownPostResponse struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IndexMarkdownPostResponse struct {
	MarkdownPosts []ShowMarkdownPostResponse `json:"markdown_posts"`
}

type CreateMarkdownPostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type UpdateMarkdownPostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
