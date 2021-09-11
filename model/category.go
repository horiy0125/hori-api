package model

import "time"

type Category struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CategoryResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type IndexCategoryResponse struct {
	Categories []CategoryResponse `json:"categories"`
}

type CategoryRequest struct {
	Name string `json:"name"`
}
