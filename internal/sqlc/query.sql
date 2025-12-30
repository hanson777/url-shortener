-- name: CreateShortURL :one
INSERT INTO urls (long_url)
VALUES ($1)
RETURNING *;

-- name: GetLongURL :one
SELECT * FROM urls
WHERE id = $1 LIMIT 1;