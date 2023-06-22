-- name: GetProduct :one
SELECT *
FROM products
WHERE id = $1
LIMIT 1;

-- name: ListProducts :many
SELECT *
FROM products
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateProduct :exec
UPDATE products
SET price  = $2,
    name   = $3,
    status = $4
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE
FROM products
WHERE id = $1;


-- name: CreateProduct :one
INSERT INTO products(id, name, merchant_id, price, status, created_at, updated_at, category_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;