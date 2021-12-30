-- DO
-- $do$
-- BEGIN
--    IF NOT EXISTS (
--       SELECT FROM pg_catalog.pg_roles  -- SELECT list can be empty for this
--       WHERE  rolname = 'payment') THEN

--       CREATE ROLE payment LOGIN PASSWORD 'payment';
--    END IF;
-- END
-- $do$;

-- CREATE DATABASE IF NOT EXISTS payment_db;


CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    iin text UNIQUE NOT NULL,
    login text UNIQUE NOT NULL,
    password text NOT NULL,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE DOMAIN role_name AS text CHECK (VALUE IN ('ROLE_ADMIN', 'ROLE_USER')); 

CREATE TABLE IF NOT EXISTS roles (
    id bigserial PRIMARY KEY,
    name role_name NOT NULL DEFAULT 'ROLE_USER',
    user_id integer UNIQUE REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS wallets (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    number bigint UNIQUE NOT NULL CHECK (number >= 111111111111 AND number <= 999999999999),
    balance double precision DEFAULT 0 CHECK (balance  >= 0),
    iin text REFERENCES users(iin),
    UNIQUE (number, iin),
    UNIQUE (iin, name)
);


-- transfer wallet from one wallet to another
-- transfer table stores: transfereriin, fromwalletid, towalletid, amount, date

CREATE TABLE IF NOT EXISTS transfers (
    id bigserial PRIMARY KEY,
    iin text REFERENCES users(iin),
    fromId bigint REFERENCES wallets(id),
    toId bigint REFERENCES wallets(id),
    amount double precision DEFAULT 0 CHECK (amount  >= 0),
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);