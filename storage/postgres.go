package storage

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joeluhrman/Lift-Tracker/types"
)

const (
	pgTableUser        = "users"
	pgTableSession     = "sessions"
	pgTableLogSetgroup = "logged_setgroups"
	pgTableLogExercise = "logged_exercises"
	pgTableLogWorkout  = "logged_workouts"
)

type PostgresStorage struct {
	conn   *sql.DB
	driver string
	url    string
}

// Call PostgresStorage.MustConnect() to initialize connection.
func NewPostgresStorage(driver string, url string) *PostgresStorage {
	return &PostgresStorage{
		driver: driver,
		url:    url,
	}
}

func (p *PostgresStorage) MustConnect() {
	var err error
	p.conn, err = sql.Open(p.driver, p.url)
	if err != nil {
		panic(err)
	}

	err = p.conn.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("connected to database %s", p.url)
}

func (p *PostgresStorage) MustClose() {
	err := p.conn.Close()
	if err != nil {
		panic(err)
	}

	log.Printf("connection to database %s successfuly close", p.url)
}

func (p *PostgresStorage) CreateUser(user *types.User, isAdmin bool) error {
	statement := "INSERT INTO " + pgTableUser + " (username, hashed_password, is_admin) VALUES ($1, $2, $3)"
	_, err := p.conn.Exec(statement, user.Username, user.HashedPassword, isAdmin)

	return err
}

func (p *PostgresStorage) AuthenticateUser(username string, password string) (int, error) {
	var (
		userID         int
		hashedPassword string
	)

	statement := "SELECT id, hashed_password FROM " + pgTableUser + " WHERE username = $1"
	row := p.conn.QueryRow(statement, username)
	if err := row.Scan(&userID, &hashedPassword); err != nil {
		return 0, err
	}

	if !passwordMatchesHash(password, hashedPassword) {
		return 0, errors.New("incorrect password")
	}

	return userID, nil
}

func (p *PostgresStorage) CreateSession(s *types.Session) error {
	statement := "INSERT INTO " + pgTableSession + " (user_id, token) VALUES ($1, $2)"
	_, err := p.conn.Exec(statement, s.UserID, s.Token)

	return err
}

func (p *PostgresStorage) AuthenticateSession(token string) (int, error) {
	var userID int

	statement := "SELECT user_id FROM " + pgTableSession + " WHERE token = $1"
	err := p.conn.QueryRow(statement, token).Scan(&userID)

	return userID, err
}

func (p *PostgresStorage) DeleteSessionByUserID(userID int) error {
	statement := "DELETE FROM " + pgTableSession + " WHERE user_id = $1"
	_, err := p.conn.Exec(statement, userID)

	return err
}

func (p *PostgresStorage) DeleteSessionByToken(token string) error {
	statement := "DELETE FROM " + pgTableSession + " WHERE token = $1"
	_, err := p.conn.Exec(statement, token)

	return err
}

/*
func (p *PostgresStorage) CreateLoggedWorkout(w *types.Workout) error {
	statement := "INSERT INTO " + pgTableLogWorkout +
		" (user_id, name, time, notes) VALUES ($1, $2, $3, $4) " +
		"RETURNING id"

	err := p.conn.QueryRow(statement, w.UserID, w.Name, w.Time, w.Notes).Scan(&w.ID)
	if err != nil {
		return err
	}

	for _, e := range w.Exercises {
		e.WorkoutID = w.ID

		statement := "INSERT INTO " + pgTableLogExercise +
			" (workout_id, name, notes) VALUES ($1, $2, $3) " +
			"RETURNING id"

		err := p.conn.QueryRow(statement, e.WorkoutID, e.Name, e.Notes).Scan(&e.ID)
		if err != nil {
			return err
		}

		for _, s := range e.Setgroups {
			s.ExerciseID = e.ID

			statement := "INSERT INTO " + pgTableLogSetgroup +
				" (exercise_id, weight, sets, reps) VALUES ($1, $2, $3, $4) " +
				"RETURNING id"

			err := p.conn.QueryRow(statement, s.ExerciseID, s.Weight, s.Sets, s.Reps).Scan(&s.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
*/
