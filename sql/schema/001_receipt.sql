-- +goose Up
CREATE TABLE
    receipts (
        id TEXT PRIMARY KEY,
        retailer TEXT NOT NULL,
        purchase_datetime DATETIME NOT NULL,
        total DECIMAL(10, 2) NOT NULL DEFAULT 0.00
    );

-- +goose Down
DROP TABLE IF EXISTS receipts;