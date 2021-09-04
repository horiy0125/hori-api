package model

import "time"

type ExternalPost struct {
	Id           int64     `db:"id"`
	Title        string    `db:"title"`
	Url          string    `db:"url"`
	ThumbnailUrl string    `db:"thumbnail_url"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	PublishedAt  time.Time `db:"published_at"`
}

type ShowExternalPostResponse struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	PublishedAt  time.Time `json:"published_at"`
}

type IndexExternalPostResponse struct {
	ExternalPosts []ShowExternalPostResponse `json:"ExternalPosts"`
}

type CreateExternalPostRequest struct {
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	PublishedAt  time.Time `json:"published_at"`
}

type UpdateExternalPostRequest struct {
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	PublishedAt  time.Time `json:"published_at"`
}
