package model

import (
	"database/sql"
	"time"
)

type MarkdownPost struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	CategoryId int64     `json:"categoryId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type NullableMarkdownPost struct {
	Id         int64         `db:"id"`
	Title      string        `db:"title"`
	Body       string        `db:"body"`
	CategoryId sql.NullInt64 `db:"category_id"`
	CreatedAt  time.Time     `db:"created_at"`
	UpdatedAt  time.Time     `db:"updated_at"`
}

type IndexMarkdownPostResponse struct {
	MarkdownPosts []MarkdownPost `json:"markdownPosts"`
}

type CreateMarkdownPostRequest struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryId int64  `json:"categoryId"`
}

type UpdateMarkdownPostRequest struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryId int64  `json:"categoryId"`
}
