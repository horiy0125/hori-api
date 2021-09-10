package usecase

import (
	"fmt"

	"github.com/horri1520/hori-api/db"
	"github.com/horri1520/hori-api/model"
	"github.com/horri1520/hori-api/repository"
	"github.com/jmoiron/sqlx"
)

type CategoryUsecase struct {
	db *sqlx.DB
}

func NewCategoryUsecase(db *sqlx.DB) *CategoryUsecase {
	return &CategoryUsecase{
		db: db,
	}
}

func (u *CategoryUsecase) Show(id int64) (*model.Category, error) {
	category, err := repository.FindCategory(u.db, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (u *CategoryUsecase) Index() ([]model.Category, error) {
	categories, err := repository.AllCategories(u.db)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (u *CategoryUsecase) Create(name string) (int64, error) {
	newCategory := model.Category{
		Name: name,
	}

	var createdId int64
	if err := db.TXHandler(u.db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertCategory(tx, newCategory)
		if err != nil {
			return err
		}

		createdId = id
		if err := tx.Commit(); err != nil {
			return err
		}

		return err
	}); err != nil {
		return 0, fmt.Errorf("failed category insert transaction: %w", err)
	}

	return createdId, nil
}

func (u *CategoryUsecase) Update(id int64, name string) error {
	updatedCategory := &model.Category{
		Id:   id,
		Name: name,
	}

	if err := db.TXHandler(u.db, func(tx *sqlx.Tx) error {
		err := repository.UpdateCategory(tx, updatedCategory)
		if err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		return err
	}); err != nil {
		return fmt.Errorf("failed category update transaction: %w", err)
	}

	return nil
}

func (u *CategoryUsecase) Destroy(requestedID int64) error {
	if err := db.TXHandler(u.db, func(tx *sqlx.Tx) error {
		err := repository.DeleteCategory(tx, requestedID)
		if err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		return err
	}); err != nil {
		return fmt.Errorf("failed category delete transaction: %w", err)
	}

	return nil
}
