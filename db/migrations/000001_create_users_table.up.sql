CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    username TEXT NOT NULL,
    avatar TEXT,
    firstname TEXT NOT NULL,
    surname TEXT,
    is_verified BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);