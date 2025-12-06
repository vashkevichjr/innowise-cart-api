-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts(
                                    id SERIAL PRIMARY KEY,

                                    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS items (
                                     id SERIAL PRIMARY KEY,
                                     product VARCHAR(255) NOT NULL,
                                     price REAL NOT NULL,

                                     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS cart_items (
                                          cart_id INT NOT NULL,
                                          item_id INT NOT NULL,
                                          quantity INT NOT NULL,

                                          PRIMARY KEY (cart_id, item_id),

                                          CONSTRAINT fk_cart
                                              FOREIGN KEY (cart_id)
                                                  REFERENCES carts(id)
                                                  ON DELETE CASCADE,

                                          CONSTRAINT fk_item
                                              FOREIGN KEY (item_id)
                                                  REFERENCES items(id)
                                                  ON DELETE CASCADE,


                                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                          deleted_at TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS items;
-- +goose StatementEnd