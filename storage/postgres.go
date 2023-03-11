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
	pgTableUser             = "users"
	pgTableSession          = "sessions"
	pgTableExerciseType     = "exercise_types"
	pgTableSetGroupLog      = "setgroup_logs"
	pgTableExerciseLog      = "exercise_logs"
	pgTableWorkoutLog       = "workout_logs"
	pgTableSetGroupTemplate = "setgroup_templates"
	pgTableExerciseTemplate = "exercise_templates"
	pgTableWorkoutTemplate  = "workout_templates"
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

	log.Printf("connection to database %s successfuly closed", p.url)
}

func (p *PostgresStorage) CreateUser(user *types.User) error {
	var err error
	user.HashedPassword, err = hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.IsAdmin = false

	statement := "INSERT INTO " + pgTableUser + " (username, hashed_password, is_admin) VALUES ($1, $2, $3) " +
		"RETURNING (id)"
	err = p.conn.QueryRow(statement, user.Username, user.HashedPassword, user.IsAdmin).Scan(&user.ID)

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

// currently just used for development, maybe later as part of an admin endpoint?
func (p *PostgresStorage) CreateExerciseType(exerciseType *types.ExerciseType) error {
	pngBytes, err := pngToBytes(exerciseType.Image)
	if err != nil {
		return err
	}

	statement := "INSERT INTO " + pgTableExerciseType + " (name, image, ppl_type, muscle_group) " +
		"VALUES ($1, $2, $3, $4) RETURNING (id)"

	err = p.conn.QueryRow(statement, exerciseType.Name, pngBytes,
		exerciseType.PPLType, exerciseType.MuscleGroup).Scan(&exerciseType.ID)

	return err
}

func (p *PostgresStorage) GetExerciseTypes() ([]types.ExerciseType, error) {
	var exerciseTypes []types.ExerciseType

	statement := "SELECT * FROM " + pgTableExerciseType

	rows, err := p.conn.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var exType types.ExerciseType
		if err := rows.Scan(exType.ID, exType.Name, exType.Image, exType.PPLType,
			exType.MuscleGroup, exType.CreatedAt, exType.UpdatedAt); err != nil {
			return nil, err
		}
		exerciseTypes = append(exerciseTypes, exType)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exerciseTypes, nil
}

func (p *PostgresStorage) CreateWorkoutTemplate(workoutTemplate *types.WorkoutTemplate) error {
	var (
		wtStatement = "INSERT INTO " + pgTableWorkoutTemplate + " (user_id, name) " +
			"VALUES ($1, $2) RETURNING id, created_at, updated_at"
		etStatement = "INSERT INTO " + pgTableExerciseTemplate + " (workout_template_id, exercise_type_id) " +
			"VALUES ($1, $2) RETURNING id, created_at, updated_at"
		sgtStatment = "INSERT INTO " + pgTableSetGroupTemplate + " (exercise_template_id, sets, reps) " +
			"VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	)

	tx, err := p.conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.QueryRow(wtStatement, workoutTemplate.UserID, workoutTemplate.Name).
		Scan(&workoutTemplate.ID, &workoutTemplate.CreatedAt, &workoutTemplate.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	for i := range workoutTemplate.ExerciseTemplates {
		workoutTemplate.ExerciseTemplates[i].WorkoutTemplateID = workoutTemplate.ID

		err = tx.QueryRow(etStatement, workoutTemplate.ID, workoutTemplate.ExerciseTemplates[i].ExerciseTypeID).
			Scan(&workoutTemplate.ExerciseTemplates[i].ID, &workoutTemplate.ExerciseTemplates[i].CreatedAt,
				&workoutTemplate.ExerciseTemplates[i].UpdatedAt)
		if err != nil {
			tx.Rollback()
			return err
		}

		for j := range workoutTemplate.ExerciseTemplates[i].SetGroupTemplates {
			workoutTemplate.ExerciseTemplates[i].SetGroupTemplates[j].ExerciseTemplateID =
				workoutTemplate.ExerciseTemplates[i].ID

			// brace for this gargantuan line
			err = tx.QueryRow(
				sgtStatment, workoutTemplate.ExerciseTemplates[i].SetGroupTemplates[j].ExerciseTemplateID,
				workoutTemplate.ExerciseTemplates[i].SetGroupTemplates[j].Sets,
				workoutTemplate.ExerciseTemplates[i].SetGroupTemplates[j].Reps).
				Scan(
					&workoutTemplate.ExerciseTemplates[i].SetGroupTemplates[j].ID,
					&workoutTemplate.ExerciseTemplates[i].SetGroupTemplates[j].CreatedAt,
					&workoutTemplate.ExerciseTemplates[i].SetGroupTemplates[j].UpdatedAt,
				)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}

func (p *PostgresStorage) GetWorkoutTemplates(userID uint) ([]types.WorkoutTemplate, error) {
	var (
		wTemps       []types.WorkoutTemplate
		wtStatement  = "SELECT * FROM " + pgTableWorkoutTemplate + " WHERE user_id = $1"
		etStatement  = "SELECT * FROM " + pgTableExerciseTemplate + " WHERE workout_template_id = $1"
		sgtStatement = "SELECT * FROM " + pgTableSetGroupTemplate + " WHERE exercise_template_id = $1"
	)

	wRows, err := p.conn.Query(wtStatement, userID)
	if err != nil {
		return nil, err
	}

	for wRows.Next() {
		var (
			w      = types.WorkoutTemplate{}
			eTemps []types.ExerciseTemplate
		)

		if err := wRows.Scan(&w.ID, &w.UserID, &w.Name, &w.CreatedAt, &w.UpdatedAt); err != nil {
			return nil, err
		}

		eRows, err := p.conn.Query(etStatement, w.ID)
		if err != nil {
			return nil, err
		}

		for eRows.Next() {
			var (
				e       = types.ExerciseTemplate{}
				sgTemps []types.SetGroupTemplate
			)

			if err := eRows.Scan(&e.ID, &e.WorkoutTemplateID, &e.ExerciseTypeID, &e.CreatedAt, &e.UpdatedAt); err != nil {
				return nil, err
			}

			sgRows, err := p.conn.Query(sgtStatement, e.ID)
			if err != nil {
				return nil, err
			}

			for sgRows.Next() {
				sgt := types.SetGroupTemplate{}
				if err := sgRows.Scan(&sgt.ID, &sgt.ExerciseTemplateID, &sgt.Sets, &sgt.Reps,
					&sgt.CreatedAt, &sgt.UpdatedAt); err != nil {
					return nil, err
				}

				sgTemps = append(sgTemps, sgt)
			}

			e.SetGroupTemplates = sgTemps
			eTemps = append(eTemps, e)
		}

		w.ExerciseTemplates = eTemps
		wTemps = append(wTemps, w)
	}

	return wTemps, nil
}
