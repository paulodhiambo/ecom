-- name: GetCountry :one
SELECT *
FROM countries
WHERE code = $1
LIMIT 1;

-- name: ListCountries :many
SELECT *
FROM countries
ORDER BY code
LIMIT $1 OFFSET $2;

-- name: UpdateCountry :exec
UPDATE countries
SET name = $2
WHERE code = $1
RETURNING *;

-- name: DeleteCountry :exec
DELETE
FROM countries
WHERE code = $1;


-- name: CreateCountry :one
INSERT INTO countries(code, name, continent_name)
VALUES ($1, $2, $3)
RETURNING *;