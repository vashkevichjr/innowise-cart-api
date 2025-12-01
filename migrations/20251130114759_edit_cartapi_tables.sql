-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts(
    id SERIAL PRIMARY KEY,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE IF NOT EXISTS items (
                                     id SERIAL PRIMARY KEY,
                                     product VARCHAR(255) NOT NULL,
                                     price DECIMAL NOT NULL,

                                     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS cart_items (
    cart_id BIGINT NOT NULL,
    item_id BIGINT NOT NULL,
    quantity INT,
    PRIMARY KEY (cart_id, item_id),

    CONSTRAINT fk_carts
    FOREIGN KEY (cart_id)
    REFERENCES carts(id)
    ON DELETE CASCADE,

    CONSTRAINT fk_items
    FOREIGN KEY (item_id)
    REFERENCES items(id)
    ON DELETE CASCADE,

    /* добавить ограничение для объединения ключей*/

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS carts;
-- +goose StatementEnd