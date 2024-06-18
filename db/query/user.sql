-- name: CreateUser :one
INSERT INTO
  "user"("username", "email", "phone_number")
VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT
  *
FROM
  "user"
WHERE
  "id" = $1;

