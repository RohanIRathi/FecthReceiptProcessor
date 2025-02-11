// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: receipt_items.sql

package database_util

import (
	"context"
)

const addItem = `-- name: AddItem :one
INSERT INTO
    items (id, description, price, receipt)
VALUES
    (?, ?, ?, ?) RETURNING id, description, price, receipt
`

type AddItemParams struct {
	ID          string
	Description string
	Price       float64
	Receipt     string
}

func (q *Queries) AddItem(ctx context.Context, arg AddItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, addItem,
		arg.ID,
		arg.Description,
		arg.Price,
		arg.Receipt,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Price,
		&i.Receipt,
	)
	return i, err
}

const getReceiptItems = `-- name: GetReceiptItems :many
SELECT
    id, description, price, receipt
FROM
    items
WHERE
    receipt = ?
`

func (q *Queries) GetReceiptItems(ctx context.Context, receipt string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getReceiptItems, receipt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.Price,
			&i.Receipt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
