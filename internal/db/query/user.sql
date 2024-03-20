-- name: CreateUser :one
INSERT INTO users (
    name,
    hashed_password,
    phone,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users 
SET 
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    phone = COALESCE(sqlc.narg(phone), phone),
    email = COALESCE(sqlc.narg(email), email),
    is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
    name = sqlc.arg(name)
RETURNING *;