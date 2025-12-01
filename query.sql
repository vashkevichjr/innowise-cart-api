-- name: CreateCart :one
INSERT INTO carts DEFAULT VALUES
RETURNING id;

-- name: ViewCarts :many
SELECT id, created_at, updated_at
FROM carts
WHERE cart_id = $1;

-- name: AddToCart :one
INSERT INTO items (cart_id, product, price)
VALUES ($1, $2, $3)
RETURNING id;


-- name: RemoveFromCart :exec
DELETE FROM cart_items
WHERE id = $1;



-- name: DeleteCart :exec
DELETE FROM carts
WHERE id = $1;

-- name: UpdateCartItem :one
UPDATE cart_items
SET product = $2, price = $3
WHERE id = $1
RETURNING id;