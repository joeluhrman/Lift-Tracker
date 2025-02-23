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
  email             TEXT UNIQUE NOT NULL,
  hashed_password   TEXT NOT NULL,
  is_admin          BOOLEAN,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS sessions (
  user_id     INTEGER PRIMARY KEY,
  token       TEXT UNIQUE,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS exercise_types (
  id                SERIAL PRIMARY KEY,
  name              TEXT UNIQUE NOT NULL,
  ppl_types         TEXT[],
  muscle_groups     TEXT[],
  is_default        BOOLEAN DEFAULT FALSE,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON exercise_types
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS workout_logs (
  id                SERIAL PRIMARY KEY,
  user_id           INTEGER ,
  date              TIMESTAMPTZ NOT NULL,
  name              TEXT NOT NULL,
  notes             TEXT,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_user
    FOREIGN KEY(user_id)
      REFERENCES users(id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout_logs
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS exercise_logs (
  id                SERIAL PRIMARY KEY,
  workout_log_id    INTEGER ,
  exercise_type_id  INTEGER ,
  notes             TEXT,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_workout_log
    FOREIGN KEY(workout_log_id)
      REFERENCES workout_logs(id),

  CONSTRAINT fk_exercise_type
    FOREIGN KEY(exercise_type_id)
      REFERENCES exercise_types(id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON exercise_logs
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS setgroup_logs (
  id                SERIAL PRIMARY KEY,
  exercise_log_id   INTEGER ,
  sets              INTEGER,
  reps              INTEGER,
  weight            DECIMAL,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_exercise_log
    FOREIGN KEY(exercise_log_id)
      REFERENCES exercise_logs(id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON setgroup_logs
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS workout_templates (
  id                    SERIAL PRIMARY KEY,
  user_id               INTEGER ,
  name                  TEXT,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_user
    FOREIGN KEY(user_id)
      REFERENCES users(id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout_templates
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS exercise_templates (
  id                    SERIAL PRIMARY KEY,
  workout_template_id   INTEGER ,
  exercise_type_id      INTEGER ,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_workout_template
    FOREIGN KEY(workout_template_id)
      REFERENCES workout_templates(id),

  CONSTRAINT fk_exercise_type
    FOREIGN KEY(exercise_type_id)
      REFERENCES exercise_types(id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON exercise_templates
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS setgroup_templates (
  id                    SERIAL PRIMARY KEY,
  exercise_template_id  INTEGER ,
  sets                  INTEGER NOT NULL,
  reps                  INTEGER NOT NULL,

  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_exercise_template
    FOREIGN KEY(exercise_template_id)
      REFERENCES exercise_templates(id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON setgroup_templates
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();