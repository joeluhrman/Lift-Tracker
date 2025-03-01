package storage

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joeluhrman/Lift-Tracker/types"
	"github.com/lib/pq"
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

type postgres struct {
	conn   *sql.DB
	driver string
	url    string
}

func NewPostgres(driver string, url string) *postgres {
	return &postgres{
		driver: driver,
		url:    url,
	}
}

func (p *postgres) MustOpen() {
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

func (p *postgres) MustClose() {
	err := p.conn.Close()
	if err != nil {
		panic(err)
	}

	log.Printf("connection to database %s successfully closed", p.url)
}

func (p *postgres) CreateUser(user *types.User, password string) error {
	if !usernameMeetsRequirements(user.Username) {
		return errors.New("username does not meet requirements")
	}

	if !passwordMeetsRequirements(password) {
		return errors.New("password does not meet requirements")
	}

	var err error
	user.HashedPassword, err = hashPassword(password)
	if err != nil {
		return err
	}

	user.IsAdmin = false

	statement := "INSERT INTO " + pgTableUser + " (username, email, hashed_password, is_admin) " +
		"VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"

	return p.conn.QueryRow(statement, user.Username, user.Email, user.HashedPassword, user.IsAdmin).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (p *postgres) GetUser(userID uint) (types.User, error) {
	var (
		statement = "SELECT id, username, email, is_admin, created_at, updated_at FROM " + pgTableUser +
			" WHERE id = $1"
		user types.User
	)

	err := p.conn.QueryRow(statement, userID).
		Scan(&user.ID, &user.Username, &user.Email, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}

func (p *postgres) AuthenticateUser(username string, password string) (uint, error) {
	var (
		userID         uint
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

func (p *postgres) CreateSession(s types.Session) error {
	statement := "INSERT INTO " + pgTableSession + " (user_id, token) VALUES ($1, $2)"
	_, err := p.conn.Exec(statement, s.UserID(), s.Token())

	return err
}

func (p *postgres) AuthenticateSession(token string) (uint, error) {
	var userID uint

	statement := "SELECT user_id FROM " + pgTableSession + " WHERE token = $1"
	err := p.conn.QueryRow(statement, token).Scan(&userID)

	return userID, err
}

func (p *postgres) DeleteSessionByUserID(userID uint) error {
	statement := "DELETE FROM " + pgTableSession + " WHERE user_id = $1"
	_, err := p.conn.Exec(statement, userID)

	return err
}

func (p *postgres) DeleteSessionByToken(token string) error {
	statement := "DELETE FROM " + pgTableSession + " WHERE token = $1"
	_, err := p.conn.Exec(statement, token)

	return err
}

func (p *postgres) CreateExerciseType(exerciseType *types.ExerciseType) error {
	statement := "INSERT INTO " + pgTableExerciseType + " (name, ppl_types, muscle_groups, is_default) " +
		"VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"

	return p.conn.QueryRow(statement, exerciseType.Name, pq.Array(exerciseType.PPLTypes), pq.Array(exerciseType.MuscleGroups), exerciseType.IsDefault).
		Scan(&exerciseType.ID, &exerciseType.CreatedAt, &exerciseType.UpdatedAt)
}

func (p *postgres) GetExerciseTypes() ([]types.ExerciseType, error) {
	var exerciseTypes []types.ExerciseType

	statement := "SELECT * FROM " + pgTableExerciseType

	rows, err := p.conn.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		strPPLTypes := []string{}
		strMscGroups := []string{}

		var exType types.ExerciseType
		if err := rows.Scan(&exType.ID, &exType.Name, pq.Array(&strPPLTypes),
			pq.Array(&strMscGroups), &exType.IsDefault, &exType.CreatedAt, &exType.UpdatedAt); err != nil {
			return nil, err
		}

		exType.PPLTypes, err = types.PPLTypesFromStrings(strPPLTypes)
		if err != nil {
			return nil, err
		}

		exType.MuscleGroups = types.MuscleGroupsFromStrings(strMscGroups)

		exerciseTypes = append(exerciseTypes, exType)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exerciseTypes, nil
}

func (p *postgres) createSetGroupTemplates(tx *sql.Tx, exTempID uint, sgTemps []types.SetGroupTemplate) error {
	statment := "INSERT INTO " + pgTableSetGroupTemplate + " (exercise_template_id, sets, reps) " +
		"VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"

	for i := range sgTemps {
		sgTemps[i].ExerciseTemplateID = exTempID

		err := tx.QueryRow(statment, sgTemps[i].ExerciseTemplateID, sgTemps[i].Sets, sgTemps[i].Reps).
			Scan(&sgTemps[i].ID, &sgTemps[i].CreatedAt, &sgTemps[i].UpdatedAt)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *postgres) createExerciseTemplates(tx *sql.Tx, wTempID uint, eTemps []types.ExerciseTemplate) error {
	statment := "INSERT INTO " + pgTableExerciseTemplate + " (workout_template_id, exercise_type_id) " +
		"VALUES ($1, $2) RETURNING id, created_at, updated_at"

	for i := range eTemps {
		eTemps[i].WorkoutTemplateID = wTempID

		err := tx.QueryRow(statment, eTemps[i].WorkoutTemplateID, eTemps[i].ExerciseTypeID).
			Scan(&eTemps[i].ID, &eTemps[i].CreatedAt, &eTemps[i].UpdatedAt)

		if err != nil {
			return err
		}

		err = p.createSetGroupTemplates(tx, eTemps[i].ID, eTemps[i].SetGroupTemplates)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *postgres) CreateWorkoutTemplate(workoutTemplate *types.WorkoutTemplate) error {
	var (
		wtStatement = "INSERT INTO " + pgTableWorkoutTemplate + " (user_id, name) " +
			"VALUES ($1, $2) RETURNING id, created_at, updated_at"
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

	err = p.createExerciseTemplates(tx, workoutTemplate.ID, workoutTemplate.ExerciseTemplates)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
	}

	return err
}

func (p *postgres) getSetGroupTemplates(exerciseTemplateID uint) ([]types.SetGroupTemplate, error) {
	var (
		statement = "SELECT * FROM " + pgTableSetGroupTemplate + " WHERE exercise_template_id = $1"
		sgTemps   []types.SetGroupTemplate
	)

	rows, err := p.conn.Query(statement, exerciseTemplateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		sgt := types.SetGroupTemplate{}
		if err := rows.Scan(&sgt.ID, &sgt.ExerciseTemplateID, &sgt.Sets, &sgt.Reps,
			&sgt.CreatedAt, &sgt.UpdatedAt); err != nil {
			return nil, err
		}

		sgTemps = append(sgTemps, sgt)
	}

	return sgTemps, rows.Err()
}

func (p *postgres) getExerciseTemplates(workoutTemplateID uint) ([]types.ExerciseTemplate, error) {
	var (
		statement = "SELECT * FROM " + pgTableExerciseTemplate + " WHERE workout_template_id = $1"
		eTemps    []types.ExerciseTemplate
	)

	rows, err := p.conn.Query(statement, workoutTemplateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			e       = types.ExerciseTemplate{}
			sgTemps []types.SetGroupTemplate
		)

		if err := rows.Scan(&e.ID, &e.WorkoutTemplateID, &e.ExerciseTypeID, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}

		sgTemps, err := p.getSetGroupTemplates(e.ID)
		if err != nil {
			return nil, err
		}

		e.SetGroupTemplates = sgTemps
		eTemps = append(eTemps, e)
	}

	return eTemps, rows.Err()
}

func (p *postgres) GetWorkoutTemplates(userID uint) ([]types.WorkoutTemplate, error) {
	var (
		wTemps    []types.WorkoutTemplate
		statement = "SELECT * FROM " + pgTableWorkoutTemplate + " WHERE user_id = $1"
	)

	rows, err := p.conn.Query(statement, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		w := types.WorkoutTemplate{}

		if err := rows.Scan(&w.ID, &w.UserID, &w.Name, &w.CreatedAt, &w.UpdatedAt); err != nil {
			return nil, err
		}

		eTemps, err := p.getExerciseTemplates(w.ID)
		if err != nil {
			return nil, err
		}

		w.ExerciseTemplates = eTemps
		wTemps = append(wTemps, w)
	}

	return wTemps, rows.Err()
}

func (p *postgres) createSetGroupLogs(tx *sql.Tx, eLogID uint, sgLogs []types.SetGroupLog) error {
	statement := "INSERT INTO " + pgTableSetGroupLog + " (exercise_log_id, sets, reps, weight) " +
		"VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"

	for i := range sgLogs {
		sg := &sgLogs[i]
		sg.ExerciseLogID = eLogID

		err := tx.QueryRow(statement, sg.ExerciseLogID, sg.Sets, sg.Reps, sg.Weight).
			Scan(&sg.ID, &sg.CreatedAt, &sg.UpdatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *postgres) createExerciseLogs(tx *sql.Tx, wLogID uint, eLogs []types.ExerciseLog) error {
	statement := "INSERT INTO " + pgTableExerciseLog + " (workout_log_id, exercise_type_id, notes) " +
		"VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"

	for i := range eLogs {
		eLog := &eLogs[i]
		eLog.WorkoutLogID = wLogID

		err := tx.QueryRow(statement, eLog.WorkoutLogID, eLog.ExerciseTypeID, eLog.Notes).
			Scan(&eLog.ID, &eLog.CreatedAt, &eLog.UpdatedAt)
		if err != nil {
			return err
		}

		if err = p.createSetGroupLogs(tx, eLog.ID, eLog.SetGroupLogs); err != nil {
			return err
		}
	}

	return nil
}

func (p *postgres) CreateWorkoutLog(wLog *types.WorkoutLog) error {
	statement := "INSERT INTO " + pgTableWorkoutLog + " (user_id, date, name, notes) " +
		"VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"

	tx, err := p.conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.QueryRow(statement, wLog.UserID, wLog.Date, wLog.Name, wLog.Notes).
		Scan(&wLog.ID, &wLog.CreatedAt, &wLog.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = p.createExerciseLogs(tx, wLog.ID, wLog.ExerciseLogs); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (p *postgres) getSetGroupLogs(eLogID uint) ([]types.SetGroupLog, error) {
	var (
		statement = "SELECT * FROM " + pgTableSetGroupLog + " WHERE exercise_log_id = $1"
		sgLogs    []types.SetGroupLog
	)

	rows, err := p.conn.Query(statement, eLogID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		sgLog := types.SetGroupLog{}
		err = rows.Scan(&sgLog.ID, &sgLog.ExerciseLogID, &sgLog.Sets, &sgLog.Reps, &sgLog.Weight,
			&sgLog.CreatedAt, &sgLog.UpdatedAt)
		if err != nil {
			return nil, err
		}

		sgLogs = append(sgLogs, sgLog)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return sgLogs, nil
}

func (p *postgres) getExerciseLogs(wLogID uint) ([]types.ExerciseLog, error) {
	var (
		statement = "SELECT * FROM " + pgTableExerciseLog + " WHERE workout_log_id = $1"
		eLogs     []types.ExerciseLog
	)

	rows, err := p.conn.Query(statement, wLogID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		eLog := types.ExerciseLog{}
		err = rows.Scan(&eLog.ID, &eLog.WorkoutLogID, &eLog.ExerciseTypeID, &eLog.Notes,
			&eLog.CreatedAt, &eLog.UpdatedAt)
		if err != nil {
			return nil, err
		}

		sgLogs, err := p.getSetGroupLogs(eLog.ID)
		if err != nil {
			return nil, err
		}

		eLog.SetGroupLogs = sgLogs
		eLogs = append(eLogs, eLog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return eLogs, nil
}

func (p *postgres) GetWorkoutLogs(userID uint) ([]types.WorkoutLog, error) {
	var (
		statement = "SELECT * FROM " + pgTableWorkoutLog + " WHERE user_id = $1"
		wLogs     []types.WorkoutLog
	)

	rows, err := p.conn.Query(statement, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		wLog := types.WorkoutLog{}
		err := rows.Scan(&wLog.ID, &wLog.UserID, &wLog.Date, &wLog.Name,
			&wLog.Notes, &wLog.CreatedAt, &wLog.UpdatedAt)
		if err != nil {
			return nil, err
		}

		eLogs, err := p.getExerciseLogs(wLog.ID)
		if err != nil {
			return nil, err
		}

		wLog.ExerciseLogs = eLogs
		wLogs = append(wLogs, wLog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return wLogs, nil
}
