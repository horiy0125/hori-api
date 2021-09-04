package usecase

import (
	"fmt"
	"time"

	"github.com/horri1520/hori-api/db"
	"github.com/horri1520/hori-api/model"
	"github.com/horri1520/hori-api/repository"
	"github.com/jmoiron/sqlx"
)

type ExternalPostUsecase struct {
	db *sqlx.DB
}

func NewExternalPostUsecase(db *sqlx.DB) *ExternalPostUsecase {
	return &ExternalPostUsecase{
		db: db,
	}
}

func (u *ExternalPostUsecase) Show(requestedId int64) (*model.ExternalPost, error) {
	externalPost, err := repository.FindExternalPost(u.db, requestedId)
	if err != nil {
		return nil, err
	}

	return externalPost, nil
}

func (u *ExternalPostUsecase) Index() ([]model.ExternalPost, error) {
	externalPosts, err := repository.AllExternalPosts(u.db)
	if err != nil {
		return nil, err
	}

	return externalPosts, nil
}

func (u *ExternalPostUsecase) Create(title string, url string, thumbnailUrl string, publishedAt time.Time) (int64, error) {
	newExternalPost := model.ExternalPost{
		Title:        title,
		Url:          url,
		ThumbnailUrl: thumbnailUrl,
		PublishedAt:  publishedAt,
	}

	var createdId int64
	if err := db.TXHandler(u.db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertExternalPost(tx, newExternalPost)
		if err != nil {
			return err
		}

		createdId = id
		if err := tx.Commit(); err != nil {
			return err
		}

		return err
	}); err != nil {
		return 0, fmt.Errorf("failed external post insert transaction: %w", err)
	}

	return createdId, nil
}

func (u *ExternalPostUsecase) Update(id int64, title string, url string, thumbnailUrl string, publishedAt time.Time) error {
	updatedExternalPost := &model.ExternalPost{
		Id:           id,
		Title:        title,
		Url:          url,
		ThumbnailUrl: thumbnailUrl,
		PublishedAt:  publishedAt,
	}

	if err := db.TXHandler(u.db, func(tx *sqlx.Tx) error {
		err := repository.UpdateExternalPost(tx, updatedExternalPost)
		if err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		return err
	}); err != nil {
		return fmt.Errorf("failed external post update transaction: %w", err)
	}

	return nil
}

func (u *ExternalPostUsecase) Destroy(requestedId int64) error {
	if err := db.TXHandler(u.db, func(tx *sqlx.Tx) error {
		err := repository.DeleteExternalPost(tx, requestedId)
		if err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		return err
	}); err != nil {
		return fmt.Errorf("failed external post delete transaction: %w", err)
	}

	return nil
}
