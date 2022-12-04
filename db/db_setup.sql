CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    username    TEXT UNIQUE,
    password    TEXT,
    is_admin    BOOLEAN
);
