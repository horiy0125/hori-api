package repository

import (
	"time"

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

func InsertMarkdownPost(db *sqlx.Tx, markdownPost model.MarkdownPost) (int64, error) {
	stmt, err := db.Preparex("insert into markdown_posts (title, body, created_at, updated_at) values ($1, $2, $3, $4) returning id")
	if err != nil {
		return 0, nil
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	err = stmt.QueryRow(markdownPost.Title, markdownPost.Body, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
