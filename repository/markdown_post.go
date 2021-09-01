package repository

import (
	"github.com/horri1520/hori-api/model"
	"github.com/jmoiron/sqlx"
)

func FindMarkdownPost(db *sqlx.DB, id int64) (*model.MarkdownPost, error) {
	var markdownPost model.MarkdownPost

	err := db.Get(&markdownPost, "select * from markdown_posts where id = $1", id)
	if err != nil {
		return nil, err
	}

	return &markdownPost, nil
}

func AllMarkdownPosts(db *sqlx.DB) ([]model.MarkdownPost, error) {
	var markdownPosts []model.MarkdownPost

	err := db.Select(&markdownPosts, "select * from markdown_posts")
	if err != nil {
		return nil, err
	}

	return markdownPosts, nil
}
