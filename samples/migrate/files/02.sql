-- name: add-column-user-email

ALTER TABLE repos ADD COLUMN email TEXT;

-- name: create-index-email

CREATE UNIQUE INDEX email_index ON users (email)
