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

-- name: UpdateShortUrl :one
UPDATE urls
SET original_url = $1
WHERE short_url = $2
returning *;

-- name: DeleteShortUrl :exec
DELETE FROM urls
WHERE short_url = $1;

-- name: GetPagedUrls :many
SELECT * FROM urls
ORDER BY id
LIMIT $1 
OFFSET $2;
