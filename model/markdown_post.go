package model

import (
	"time"
)

type MarkdownPost struct {
	Id           int64     `db:"id"`
	Title        string    `db:"title"`
	Body         string    `db:"body"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	CategoryId   int64     `db:"category_id"`
	CategoryName string    `db:"category_name"`
	Publish      bool      `db:"publish"`
}

type MarkdownPostResponse struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CategoryId   int64     `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	Publish      bool      `json:"publish"`
}

type IndexMarkdownPostResponse struct {
	MarkdownPosts []MarkdownPostResponse `json:"markdownPosts"`
}

type MarkdownPostRequest struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryId int64  `json:"categoryId"`
	Publish    bool   `json:"publish"`
}
