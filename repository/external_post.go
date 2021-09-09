package repository

import (
	"time"

	"github.com/horri1520/hori-api/model"
	"github.com/jmoiron/sqlx"
)

func FindExternalPost(db *sqlx.DB, id int64) (*model.ExternalPost, error) {
	var nullableExternalPost model.NullableExternalPost

	err := db.Get(&nullableExternalPost, "select * from external_posts where id = $1", id)
	if err != nil {
		return nil, err
	}

	categoryId := int64(nullableExternalPost.CategoryId.Int64)
	externalPost := model.ExternalPost{
		Id:           nullableExternalPost.Id,
		Title:        nullableExternalPost.Title,
		Url:          nullableExternalPost.Url,
		ThumbnailUrl: nullableExternalPost.ThumbnailUrl,
		CategoryId:   categoryId,
		CreatedAt:    nullableExternalPost.CreatedAt,
		UpdatedAt:    nullableExternalPost.CreatedAt,
		PublishedAt:  nullableExternalPost.PublishedAt,
	}

	return &externalPost, nil
}

func AllExternalPosts(db *sqlx.DB) ([]model.ExternalPost, error) {
	var nullableExternalPosts []model.NullableExternalPost

	err := db.Select(&nullableExternalPosts, "select * from external_posts order by updated_at desc")
	if err != nil {
		return nil, err
	}

	var externalPosts []model.ExternalPost
	for _, e := range nullableExternalPosts {
		categoryId := int64(e.CategoryId.Int64)
		externalPost := model.ExternalPost{
			Id:           e.Id,
			Title:        e.Title,
			Url:          e.Url,
			ThumbnailUrl: e.ThumbnailUrl,
			CategoryId:   categoryId,
			CreatedAt:    e.CreatedAt,
			UpdatedAt:    e.UpdatedAt,
			PublishedAt:  e.PublishedAt,
		}
		externalPosts = append(externalPosts, externalPost)
	}

	return externalPosts, nil
}

func InsertExternalPost(db *sqlx.Tx, externalPost model.ExternalPost) (int64, error) {
	stmt, err := db.Preparex("insert into external_posts (title, url, thumbnail_url, category_id, created_at, updated_at, published_at) values ($1, $2, $3, $4, $5, $6, $7) returning id")
	if err != nil {
		return 0, nil
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	err = stmt.QueryRow(externalPost.Title, externalPost.Url, externalPost.ThumbnailUrl, externalPost.CategoryId, time.Now(), time.Now(), externalPost.PublishedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateExternalPost(db *sqlx.Tx, externalPost *model.ExternalPost) error {
	stmt, err := db.Preparex("update external_posts set title = $1, url = $2, thumbnail_url = $3, category_id = $4, updated_at = $5, published_at = $6 where id = $7")
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	_, err = stmt.Exec(externalPost.Title, externalPost.Url, externalPost.ThumbnailUrl, externalPost.CategoryId, time.Now(), externalPost.PublishedAt, externalPost.Id)
	if err != nil {
		return err
	}

	return nil

}

func DeleteExternalPost(db *sqlx.Tx, id int64) error {
	stmt, err := db.Preparex("delete from external_posts where id = $1")
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
