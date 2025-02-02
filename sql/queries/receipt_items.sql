-- name: AddItem :one

INSERT INTO items (id, description, price, receipt) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetReceiptItems :many
SELECT * FROM items WHERE receipt=$1;