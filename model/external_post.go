package model

import (
	"time"
)

type ExternalPost struct {
	Id           int64     `db:"id"`
	Title        string    `db:"title"`
	Url          string    `db:"url"`
	ThumbnailUrl string    `db:"thumbnail_url"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	PublishedAt  time.Time `db:"published_at"`
	CategoryId   int64     `db:"category_id"`
	CategoryName string    `db:"category_name"`
}

type ExternalPostResponse struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	PublishedAt  time.Time `json:"publishedAt"`
	CategoryId   int64     `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
}

type IndexExternalPostResponse struct {
	ExternalPosts []ExternalPostResponse `json:"externalPosts"`
}

type ExternalPostRequest struct {
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	CategoryId   int64     `json:"categoryId"`
	PublishedAt  time.Time `json:"publishedAt"`
}
