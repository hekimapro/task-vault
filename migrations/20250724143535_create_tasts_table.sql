-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    tasts (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        user_id UUID REFERENCES users (id) ON DELETE CASCADE NOT NULL,
        title VARCHAR(100) NOT NULL,
        description TEXT DEFAULT NULL,
        status TEXT CHECK (status IN ('pending', 'done')) DEFAULT 'pending',
        due_date DATE,
        created_at TIMESTAMPTZ DEFAULT NOW (),
        updated_at TIMESTAMPTZ DEFAULT NOW ()
    );

-- +goose Down
DROP TABLE IF EXISTS tasts CASCADE;