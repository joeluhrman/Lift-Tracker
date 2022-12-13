CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    username    TEXT UNIQUE NOT NULL,
    hashed_password    TEXT NOT NULL,
    is_admin    BOOLEAN,

    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sessions (
    user_id     INTEGER PRIMARY KEY,
    token       TEXT UNIQUE,

    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);