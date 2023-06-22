-- name: GetMerchant :one
SELECT *
FROM merchants
WHERE id = $1
LIMIT 1;

-- name: ListMerchants :many
SELECT *
FROM merchants
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateMerchant :exec
UPDATE merchants
SET merchant_name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteMerchant :exec
DELETE
FROM merchants
WHERE id = $1;


-- name: CreateMerchant :one
INSERT INTO merchants(id, merchant_name, country_code, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;