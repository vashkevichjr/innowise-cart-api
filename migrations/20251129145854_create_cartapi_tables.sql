-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts(
                       id SERIAL primary key
);

CREATE TABLE IF NOT EXISTS cart_items (
                         id SERIAL primary key,
                         cart_id BIGINT NOT NULL ,
                         product varchar(255) NOT NULL ,
                         price BIGINT NOT NULL,

                         CONSTRAINT fk_carts
                             FOREIGN KEY (cart_id)
                                 REFERENCES carts(id)
                                 ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;
-- +goose StatementEnd