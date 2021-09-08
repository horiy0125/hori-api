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

	err := db.Select(&markdownPosts, "select * from markdown_posts order by updated_at desc")
	if err != nil {
		return nil, err
	}

	return markdownPosts, nil
}

func InsertMarkdownPost(db *sqlx.Tx, markdownPost model.MarkdownPost) (int64, error) {
	stmt, err := db.Preparex("insert into markdown_posts (title, body, category_id, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id")
	if err != nil {
		return 0, nil
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	err = stmt.QueryRow(markdownPost.Title, markdownPost.Body, markdownPost.CategoryId, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateMarkdownPost(db *sqlx.Tx, markdownPost *model.MarkdownPost) error {
	stmt, err := db.Preparex("update markdown_posts set title = $1, body = $2, category_id = $3, updated_at = $4 where id = $5")
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	_, err = stmt.Exec(markdownPost.Title, markdownPost.Body, markdownPost.CategoryId, time.Now(), markdownPost.Id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMarkdownPost(db *sqlx.Tx, id int64) error {
	stmt, err := db.Preparex("delete from markdown_posts where id = $1")
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
