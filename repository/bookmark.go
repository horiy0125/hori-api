package repository

import (
	"time"

	"github.com/horri1520/hori-api/model"
	"github.com/jmoiron/sqlx"
)

func FindBookmark(db *sqlx.DB, id int64) (*model.Bookmark, error) {
	var bookmark model.Bookmark

	err := db.Get(&bookmark, "select * from bookmarks where id = $1", id)
	if err != nil {
		return nil, err
	}

	return &bookmark, nil
}

func AllBookmarks(db *sqlx.DB) ([]model.Bookmark, error) {
	var bookmarks []model.Bookmark

	err := db.Select(&bookmarks, "select * from bookmarks order by updated_at desc")
	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func InsertBookmark(db *sqlx.Tx, bookmark model.Bookmark) (int64, error) {
	stmt, err := db.Preparex("insert into bookmarks (url, description, created_at, updated_at) values ($1, $2, $3, $4) returning id")
	if err != nil {
		return 0, nil
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	err = stmt.QueryRow(bookmark.Url, bookmark.Description, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
