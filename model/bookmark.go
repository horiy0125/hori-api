package model

import "time"

type Bookmark struct {
	Id           int64     `db:"id"`
	Url          string    `db:"url"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	CategoryId   int64     `db:"category_id"`
	CategoryName string    `db:"category_name"`
}

type BookmarkResponse struct {
	Id           int64     `json:"id"`
	Url          string    `json:"url"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CategoryId   int64     `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
}

type IndexBookmarkResponse struct {
	Bookmarks []BookmarkResponse `json:"bookmarks"`
}

type BookmarkRequest struct {
	Url         string `json:"url"`
	Description string `json:"description"`
	CategoryId  int64  `json:"categoryId"`
}
