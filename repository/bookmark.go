package repository

import (
	"time"

	"github.com/horri1520/hori-api/model"
	"github.com/jmoiron/sqlx"
)

func FindBookmark(db *sqlx.DB, id int64) (*model.Bookmark, error) {
	var bookmark model.Bookmark

	err := db.Get(&bookmark, "select b.id, b.url, b.description, b.created_at, b.updated_at, c.id as category_id, c.name as category_name from bookmarks as b join categories as c on b.category_id = c.id where b.id = $1", id)
	if err != nil {
		return nil, err
	}

	return &bookmark, nil
}

func AllBookmarks(db *sqlx.DB) ([]model.Bookmark, error) {
	var bookmarks []model.Bookmark

	err := db.Select(&bookmarks, "select b.id, b.url, b.description, b.created_at, b.updated_at, c.id as category_id, c.name as category_name from bookmarks as b join categories as c on b.category_id = c.id order by updated_at desc")
	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func InsertBookmark(db *sqlx.Tx, bookmark model.Bookmark) (int64, error) {
	stmt, err := db.Preparex("insert into bookmarks (url, description, category_id, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id")
	if err != nil {
		return 0, nil
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	err = stmt.QueryRow(bookmark.Url, bookmark.Description, bookmark.CategoryId, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateBookmark(db *sqlx.Tx, bookmark *model.Bookmark) error {
	stmt, err := db.Preparex("update bookmarks set url = $1, description = $2, category_id = $3, updated_at = $4 where id = $5")
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	_, err = stmt.Exec(bookmark.Url, bookmark.Description, bookmark.CategoryId, time.Now(), bookmark.Id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBookmark(db *sqlx.Tx, id int64) error {
	stmt, err := db.Preparex("delete from bookmarks where id = $1")
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
