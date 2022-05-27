CREATE TYPE taxi_types AS ENUM ('economy', 'comfort', 'business', 'electro');

CREATE TABLE IF NOT EXISTS drivers
(
    id SERIAL PRIMARY KEY,
    driver_uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(25) UNIQUE NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    taxi_type taxi_types NOT NULL,
    is_busy BOOLEAN NOT NULL DEFAULT FALSE,
    driver_rating REAL NOT NULL DEFAULT 0.0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- name: CreateDriver :one
INSERT INTO drivers (name, phone_number, email, password, taxi_type)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetDrivers :many
SELECT * FROM drivers;

-- name: GetDriverByID :one
SELECT * FROM drivers
WHERE id = $1;

-- name: GetDriverStatusByID :one
SELECT is_busy FROM drivers
WHERE id = $1;

-- name: GetDriverPhoneByID :one
SELECT phone_number FROM drivers
WHERE id = $1;

-- name: GetDriverEmailByID :one
SELECT email FROM drivers
WHERE id = $1;

-- name: GetDriverRatingByID :one
SELECT driver_rating FROM drivers
WHERE id = $1;

-- name: UpdateDriver :one
UPDATE drivers
SET name = $1, phone_number = $2, email = $3, password = $4, taxi_type = $5
WHERE id=$6 RETURNING *;

-- name: DeleteDriver :exec
DELETE FROM drivers
WHERE id = $1;
