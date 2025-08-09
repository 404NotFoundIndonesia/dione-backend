CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(255),
    email       VARCHAR(255) UNIQUE,
    phone       VARCHAR(50),
    role        VARCHAR(50) CHECK (role IN ('user', 'admin')),
    password    VARCHAR(255),
    bio         TEXT,
    avatar_path VARCHAR(255),
    is_active   BOOLEAN          DEFAULT TRUE,
    created_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);