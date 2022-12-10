CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    username    TEXT UNIQUE NOT NULL,
    hashed_password    TEXT NOT NULL,
    is_admin    BOOLEAN
);

CREATE TABLE IF NOT EXISTS sessions (
    user_id     INTEGER PRIMARY KEY,
    token       TEXT UNIQUE
);