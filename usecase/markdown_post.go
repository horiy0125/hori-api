package usecase

import (
	"fmt"

	"github.com/horri1520/hori-api/db"
	"github.com/horri1520/hori-api/model"
	"github.com/horri1520/hori-api/repository"
	"github.com/jmoiron/sqlx"
)

type MarkdownPostUsecase struct {
	db *sqlx.DB
}

func NewMarkdownPostUsecase(db *sqlx.DB) *MarkdownPostUsecase {
	return &MarkdownPostUsecase{
		db: db,
	}
}

func (u *MarkdownPostUsecase) Show(requestedId int64) (*model.MarkdownPost, error) {
	markdownPost, err := repository.FindMarkdownPost(u.db, requestedId)
	if err != nil {
		return nil, err
	}

	return markdownPost, nil
}

func (u *MarkdownPostUsecase) Index() ([]model.MarkdownPost, error) {
	markdownPosts, err := repository.AllMarkdownPosts(u.db)
	if err != nil {
		return nil, err
	}

	return markdownPosts, nil
}

func (u *MarkdownPostUsecase) Create(title string, body string) (int64, error) {
	newMarkdownPost := model.MarkdownPost{
		Title: title,
		Body:  body,
	}

	var createdId int64
	if err := db.TXHandler(u.db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertMarkdownPost(tx, newMarkdownPost)
		if err != nil {
			return err
		}

		createdId = id
		if err := tx.Commit(); err != nil {
			return err
		}

		return err
	}); err != nil {
		return 0, fmt.Errorf("failed markdown post insert transaction: %w", err)
	}

	return createdId, nil
}
