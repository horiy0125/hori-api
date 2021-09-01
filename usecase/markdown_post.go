package usecase

import (
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

func (u *MarkdownPostUsecase) Show(showMarkdownPostId int64) (*model.MarkdownPost, error) {
	markdownPost, err := repository.FindMarkdownPost(u.db, showMarkdownPostId)
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
