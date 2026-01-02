-- name: CreateShortURL :one
INSERT INTO urls (long_url)
VALUES ($1)
RETURNING *;

-- name: GetLongURL :one
SELECT * FROM urls
WHERE id = $1 LIMIT 1;

-- name: IncrementClicks :exec
UPDATE urls SET clicks = clicks + 1 WHERE id = $1;