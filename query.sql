-- CARTS
-- name: CreateCart :one
INSERT INTO carts DEFAULT VALUES
RETURNING id, created_at, updated_at, deleted_at;

-- name: GetCart :one
SELECT id, created_at, updated_at, deleted_at
FROM carts
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetCarts :many
SELECT id, created_at, updated_at, deleted_at
FROM carts
WHERE deleted_at IS NULL;

-- name: SoftDeleteCart :exec
UPDATE carts
SET deleted_at = NOW()
WHERE id = $1;

-- name: HardDeleteCart :exec
DELETE FROM carts
WHERE id = $1;



-- ITEMS
-- name: CreateItem :one
INSERT INTO items(product, price)
VALUES ($1,$2)
RETURNING id, product, price, created_at, updated_at, deleted_at;

-- name: UpdateItem :one
UPDATE items
SET product = $2, price = $3
WHERE id = $1
RETURNING id;

-- name: UpdateItemProduct :one
UPDATE items
SET product = $2
WHERE id = $1
RETURNING id;

-- name: UpdateItemPrice :one
UPDATE items
SET price = $2
WHERE id = $1
RETURNING id;

-- name: GetItem :one
SELECT id, product, price, created_at, updated_at
FROM items
WHERE id = $1;

-- name: GetItems :many
SELECT id, product, price, created_at, updated_at
FROM items;

-- name: SoftDeleteItem :exec
UPDATE items
SET deleted_at = NOW()
WHERE id = $1;

-- name: HardDeleteItem :exec
DELETE FROM items
WHERE id = $1;



-- CART_ITEMS
-- name: AddItemToCart :exec
INSERT INTO cart_items (cart_id, item_id, quantity)
VALUES ($1, $2, $3);

-- name: UpdateItemInCart :exec
UPDATE cart_items
SET item_id = $3, quantity = $4
WHERE cart_id = $1 and item_id = $2;
UPDATE carts
SET updated_at = NOW()
WHERE id = $1;

-- name: UpdateItemInCartQuantity :exec
UPDATE cart_items
SET quantity = $3
WHERE cart_id = $1 and item_id = $2;
UPDATE carts
SET updated_at = NOW()
WHERE id = $1;

-- name: GetItemsByCart :many
SELECT ci.cart_id, ci.item_id, ci.quantity, i.product, i.price
FROM carts c
JOIN cart_items ci ON ci.cart_id = c.id
JOIN items i ON i.id = ci.item_id
WHERE c.id = $1 and i.deleted_at IS NULL;

-- name: GetCartsByItem :many
SELECT ci.cart_id, ci.item_id, ci.quantity, i.product, i.price
FROM carts c
         JOIN cart_items ci ON ci.cart_id = c.id
         JOIN items i ON i.id = ci.cart_id
WHERE c.id = $1;


-- name: GetCartsItems :many
SELECT ci.cart_id, ci.item_id, ci.quantity, i.product, i.price
FROM carts c
         JOIN cart_items ci ON ci.cart_id = c.id
         JOIN items i ON i.id = ci.cart_id;

-- name: SoftDeleteItemByCart :exec
UPDATE cart_items
SET deleted_at = NOW()
WHERE cart_id = $1 and item_id = $2;

-- name: HardDeleteItemByCart :exec
DELETE FROM cart_items
WHERE cart_id = $1 and item_id = $2;


