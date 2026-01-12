-- name: CreateShortURL :one
INSERT INTO urls (long_url)
VALUES ($1)
RETURNING *;

-- name: GetLongURL :one
SELECT * FROM urls
WHERE id = $1 LIMIT 1;

-- name: IncrementClicks :exec
UPDATE urls SET clicks = clicks + 1 WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING id, email, created_at;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, created_at
FROM users
WHERE email = $1;
