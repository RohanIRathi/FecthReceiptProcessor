-- +goose Up
CREATE TABLE
    receipts (
        id UUID PRIMARY KEY,
        retailer TEXT NOT NULL,
        purchase_datetime TIMESTAMP NOT NULL,
        total DECIMAL(10, 2) NOT NULL DEFAULT 0.00
    );

-- +goose Down
DROP TABLE receipts;