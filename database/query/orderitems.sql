-- name: GetOrderItem :one
SELECT *
FROM order_items
WHERE order_id = $1
LIMIT 1;

-- name: ListOrderItems :many
SELECT *
FROM order_items
ORDER BY order_id
LIMIT $1 OFFSET $2;

-- name: UpdateOrderItem :exec
UPDATE order_items
SET quantity = $2
WHERE order_id = $1
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE
FROM order_items
WHERE order_id = $1;


-- name: CreateOrderItem :one
INSERT INTO order_items(order_id, product_id, quantity)
VALUES ($1, $2, $3)
RETURNING *;