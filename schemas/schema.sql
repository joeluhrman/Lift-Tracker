CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS users (
  id                SERIAL PRIMARY KEY,
  username          TEXT UNIQUE NOT NULL,
  hashed_password   TEXT NOT NULL,
  is_admin          BOOLEAN,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sessions (
  user_id     INTEGER PRIMARY KEY,
  token       TEXT UNIQUE,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS setgroups (
  id            SERIAL PRIMARY KEY,
  exercise_id   INTEGER NOT NULL,
  weight        INTEGER,
  sets          INTEGER,
  reps          INTEGER,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS exercises (
  id            SERIAL PRIMARY KEY,
  workout_id    INTEGER NOT NULL,
  name          TEXT,
  notes         TEXT, 

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS workouts (
  id            SERIAL PRIMARY KEY,
  user_id       INTEGER NOT NULL,
  name          TEXT,
  time          TIMESTAMPTZ,
  notes         TEXT,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS default_exercises (
  id            SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS custom_exercises (
  id            SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS custom_workouts (
  id            SERIAL PRIMARY KEY
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON setgroups
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON exercises
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workouts
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();