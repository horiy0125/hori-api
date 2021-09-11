package repository

import (
	"time"

	"github.com/horri1520/hori-api/model"
	"github.com/jmoiron/sqlx"
)

func FindCategory(db *sqlx.DB, id int64) (*model.Category, error) {
	var category model.Category

	err := db.Get(&category, "select * from categories where id = $1", id)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func AllCategories(db *sqlx.DB) ([]model.Category, error) {
	var categories []model.Category

	err := db.Select(&categories, "select * from categories order by id desc")
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func InsertCategory(db *sqlx.Tx, category model.Category) (int64, error) {
	stmt, err := db.Preparex("insert into categories (name, created_at, updated_at) values ($1, $2, $3) returning id")
	if err != nil {
		return 0, nil
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	err = stmt.QueryRow(category.Name, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateCategory(db *sqlx.Tx, category *model.Category) error {
	stmt, err := db.Preparex("update categories set name = $1, updated_at = $2 where id = $3")

	if err != nil {
		return err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	_, err = stmt.Exec(category.Name, time.Now(), category.Id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(db *sqlx.Tx, id int64) error {
	stmt, err := db.Preparex("delete from categories where id = $1")
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
