-- name: AddItem :one
INSERT INTO
    items (id, description, price, receipt)
VALUES
    (?, ?, ?, ?) RETURNING *;

-- name: GetReceiptItems :many
SELECT
    *
FROM
    items
WHERE
    receipt = ?;