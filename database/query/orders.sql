-- name: GetOrder :one
SELECT *
FROM orders
WHERE id = $1
LIMIT 1;

-- name: ListOrders :many
SELECT *
FROM orders
ORDER BY updated_at
LIMIT $1 OFFSET $2;

-- name: UpdateOrder :exec
UPDATE orders
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE
FROM orders
WHERE id = $1;


-- name: CreateOrder :one
INSERT INTO orders(id, user_id, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;