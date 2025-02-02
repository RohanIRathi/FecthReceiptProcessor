-- +goose Up
CREATE TABLE
    items (
        id UUID PRIMARY KEY,
        description TEXT NOT NULL,
        price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
        receipt UUID NOT NULL REFERENCES receipts (id) ON DELETE CASCADE
    );

-- +goose Down
DROP TABLE items;