// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: products.sql

package database

import (
	"context"
	"time"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products(id, name, merchant_id, price, status, created_at, updated_at, category_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, name, merchant_id, price, status, created_at, updated_at, category_id
`

type CreateProductParams struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	MerchantID int32     `json:"merchant_id"`
	Price      int32     `json:"price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CategoryID int32     `json:"category_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ID,
		arg.Name,
		arg.MerchantID,
		arg.Price,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CategoryID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.MerchantID,
		&i.Price,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CategoryID,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE
FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, merchant_id, price, status, created_at, updated_at, category_id
FROM products
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.MerchantID,
		&i.Price,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CategoryID,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, merchant_id, price, status, created_at, updated_at, category_id
FROM products
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.MerchantID,
			&i.Price,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CategoryID,
		); err != nil {
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

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET price  = $2,
    name   = $3,
    status = $4
WHERE id = $1
RETURNING id, name, merchant_id, price, status, created_at, updated_at, category_id
`

type UpdateProductParams struct {
	ID     int64  `json:"id"`
	Price  int32  `json:"price"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.ID,
		arg.Price,
		arg.Name,
		arg.Status,
	)
	return err
}
