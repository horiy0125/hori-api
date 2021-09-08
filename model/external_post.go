package model

import "time"

type ExternalPost struct {
	Id           int64     `db:"id"`
	Title        string    `db:"title"`
	Url          string    `db:"url"`
	ThumbnailUrl string    `db:"thumbnail_url"`
	CategoryId   int64     `db:"category_id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	PublishedAt  time.Time `db:"published_at"`
}

type ShowExternalPostResponse struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	CategoryId   int64     `json:"categoryId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	PublishedAt  time.Time `json:"publishedAt"`
}

type IndexExternalPostResponse struct {
	ExternalPosts []ShowExternalPostResponse `json:"externalPosts"`
}

type CreateExternalPostRequest struct {
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	CategoryId   int64     `json:"categoryId"`
	PublishedAt  time.Time `json:"publishedAt"`
}

type UpdateExternalPostRequest struct {
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	CategoryId   int64     `json:"categoryId"`
	PublishedAt  time.Time `json:"publishedAt"`
}
