-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE taxi_type AS ENUM ('economy', 'comfort', 'business', 'electro');

CREATE TABLE IF NOT EXISTS drivers
(
    id SERIAL PRIMARY KEY,
    driver_uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(25) UNIQUE NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    taxi_type taxi_type NOT NULL,
    is_busy BOOLEAN NOT NULL DEFAULT FALSE,
    driver_rating REAL NOT NULL DEFAULT 0.0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    AFTER UPDATE ON drivers
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS drivers;
DROP TYPE taxi_type;
DROP EXTENSION IF EXISTS "uuid-ossp";
DROP FUNCTION IF EXISTS trigger_set_timestamp();
DROP TRIGGER IF EXISTS set_timestamp on drivers;
-- +goose StatementEnd
