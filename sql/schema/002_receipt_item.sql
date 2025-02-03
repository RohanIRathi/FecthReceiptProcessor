-- +goose Up
CREATE TABLE
    items (
        id TEXT PRIMARY KEY,
        description TEXT NOT NULL,
        price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
        receipt TEXT NOT NULL,
        FOREIGN KEY (receipt) REFERENCES receipts (id) ON DELETE CASCADE
    );

-- +goose Down
DROP TABLE IF EXISTS items;