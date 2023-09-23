-- Define tables for sellers, buyers, products, and orders
-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- Sellers table to store seller profiles
CREATE TABLE sellers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

-- Products table to store product listings
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    seller_id INT NOT NULL REFERENCES sellers(id),
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- Add any additional product attributes here
);

-- Buyers table to store buyer profiles
CREATE TABLE buyers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

-- Orders table to store buyer-seller transactions
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    buyer_id INT NOT NULL REFERENCES buyers(id),
    product_id INT NOT NULL REFERENCES products(id),
    quantity INT NOT NULL,
    order_status VARCHAR(20) NOT NULL DEFAULT 'Pending', -- Status can be 'Pending', 'Accepted', 'Completed', etc.
);


CREATE INDEX idx_product_name ON products(name);
CREATE INDEX idx_name_price_created_at ON products (name, price, created_at);


-- +goose Down
-- SQL in this section is executed when the migration is applied.

DROP TABLE If exists sellers;

DROP TABLE If exists products;

DROP TABLE If exists buyers;

DROP TABLE If exists orders;
