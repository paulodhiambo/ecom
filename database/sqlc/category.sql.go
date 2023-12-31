// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: category.sql

package database

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories(id, cat_name)
VALUES ($1, $2)
RETURNING id, cat_name
`

type CreateCategoryParams struct {
	ID      int64  `json:"id"`
	CatName string `json:"cat_name"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.ID, arg.CatName)
	var i Category
	err := row.Scan(&i.ID, &i.CatName)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE
FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, cat_name
FROM categories
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.CatName)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, cat_name
FROM categories
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.CatName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
SET cat_name = $2
WHERE id = $1
RETURNING id, cat_name
`

type UpdateCategoryParams struct {
	ID      int64  `json:"id"`
	CatName string `json:"cat_name"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.ID, arg.CatName)
	return err
}
