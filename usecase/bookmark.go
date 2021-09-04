package usecase

import (
	"fmt"

	"github.com/horri1520/hori-api/db"
	"github.com/horri1520/hori-api/model"
	"github.com/horri1520/hori-api/repository"
	"github.com/jmoiron/sqlx"
)

type BookmarkUsecase struct {
	db *sqlx.DB
}

func NewBookmarkUsecase(db *sqlx.DB) *BookmarkUsecase {
	return &BookmarkUsecase{
		db: db,
	}
}

func (u *BookmarkUsecase) Index() ([]model.Bookmark, error) {
	bookmarks, err := repository.AllBookmarks(u.db)
	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (u *BookmarkUsecase) Create(url string, description string) (int64, error) {
	newBookmark := model.Bookmark{
		Url:         url,
		Description: description,
	}

	var createdId int64
	if err := db.TXHandler(u.db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertBookmark(tx, newBookmark)
		if err != nil {
			return err
		}

		createdId = id
		if err := tx.Commit(); err != nil {
			return err
		}

		return err
	}); err != nil {
		return 0, fmt.Errorf("failed bookmark insert transaction: %w", err)
	}

	return createdId, nil
}
