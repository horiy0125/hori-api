package model

import "time"

type Bookmark struct {
	Id          int64     `db:"id"`
	Url         string    `db:"url"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ShowBookmarkResponse struct {
	Id          int64     `json:"id"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type IndexBookmarkResponse struct {
	Bookmarks []ShowBookmarkResponse `json:"bookmarks"`
}

type CreateBookmarkRequest struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}
