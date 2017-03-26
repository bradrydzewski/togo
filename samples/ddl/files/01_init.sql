-- name: create-table-user

CREATE TABLE users (
  username TEXT
  password TEXT
)

-- name: create-index-username

CREATE UNIQUE INDEX username_index ON users (username)
