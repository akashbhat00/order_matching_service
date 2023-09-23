-- +goose Up
ALTER TABLE orders
ADD created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- +goose Down
DROP TABLE If exists orders;