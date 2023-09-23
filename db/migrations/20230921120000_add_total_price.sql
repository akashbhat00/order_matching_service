-- +goose Up
ALTER TABLE orders
ADD total_price DECIMAL(10, 2) NOT NULL;

-- +goose Down
DROP TABLE If exists orders;