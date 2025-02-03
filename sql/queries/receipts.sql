-- name: CreateReceipt :one
INSERT INTO
    receipts (id, retailer, purchase_datetime, total)
VALUES
    (?, ?, ?, ?) RETURNING *;

-- name: GetReceipt :one
SELECT
    *
FROM
    receipts
WHERE
    id = ?
LIMIT
    1;