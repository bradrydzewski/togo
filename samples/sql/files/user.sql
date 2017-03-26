-- name: user-select-all

SELECT
 username
,password
,email
FROM users

-- name: user-select-username

SELECT
 username
,password
,email
FROM users
WHERE username = ?
