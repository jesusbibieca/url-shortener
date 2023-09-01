-- name: CreateShortUrl :one
INSERT INTO urls (
  user_id,
  original_url
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetShortUrl :one
SELECT * FROM urls
WHERE short_url = $1 LIMIT 1;

-- name: GetShortUrlByID :one
SELECT * FROM urls
WHERE id = $1 LIMIT 1;
