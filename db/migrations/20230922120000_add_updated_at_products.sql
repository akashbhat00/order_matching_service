-- +goose Up
ALTER TABLE products
ADD updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- +goose Down
DROP TABLE If exists products;