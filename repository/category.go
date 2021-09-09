package repository

import (
	"github.com/horri1520/hori-api/model"
	"github.com/jmoiron/sqlx"
)

func FindCategory(db *sqlx.DB, id int64) (*model.Category, error) {
	var nullableCategory model.NullableCategory

	err := db.Get(&nullableCategory, "select * from categories where id = $1", id)
	if err != nil {
		return nil, err
	}

	category := model.Category(nullableCategory)

	return &category, nil
}
