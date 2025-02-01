-- name: CreateReceipt :one

INSERT INTO receipts
    (id, retailer, purchase_datetime, total)
VALUES
    ($1, $2, $3, $4)
RETURNING *;

-- name: GetReceipt :one
SELECT * FROM receipts where id=$1 LIMIT 1;