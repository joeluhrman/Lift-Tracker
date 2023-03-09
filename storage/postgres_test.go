package storage

import (
	"fmt"
	"image"
	"os"
	"testing"

	"github.com/joeluhrman/Lift-Tracker/types"
)

var (
	testPGDriver = "pgx"
	testPGApiKey = string(MustReadFile("../api_keys/api_key_test.txt"))
	testPGURL    = "postgresql://jaluhrman:" + testPGApiKey + "@db.bit.io/jaluhrman/Lift-Tracker-Test"

	testPGStorage = &testPostgresStorage{
		PostgresStorage: NewPostgresStorage(testPGDriver, testPGURL),
	}

	tables = []string{pgTableUser, pgTableSession, pgTableExerciseType,
		pgTableSetGroupLog, pgTableExerciseLog, pgTableWorkoutLog,
		pgTableSetGroupTemplate, pgTableExerciseTemplate, pgTableWorkoutTemplate}
)

// wrapper for test methods to avoid confusion
type testPostgresStorage struct {
	*PostgresStorage
}

func (t *testPostgresStorage) clearAllTables() {
	for _, table := range tables {
		t.clearTable(table)
	}
}

func (t *testPostgresStorage) clearTable(tName string) {
	_, err := t.conn.Exec("DELETE FROM " + tName)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	testPGStorage.MustConnect()
	testPGStorage.clearAllTables()
	code := m.Run()
	os.Exit(code)
}

func Test_CreateUser(t *testing.T) {
	defer testPGStorage.clearTable(pgTableUser)

	// Success case
	func() {
		user := types.NewUser("jaluhrman", "123")

		err := testPGStorage.CreateUser(user)
		if err != nil {
			t.Error(err)
		}

		if user.ID == 0 {
			t.Error("Struct's ID was not updated")
		}
	}()

	// username already exists
	func() {
		user := types.NewUser("jaluhrman", "123")

		err := testPGStorage.CreateUser(user)
		if err == nil {
			t.Error("should have produced an error when username already taken")
		}
	}()
}

func Test_AuthenticateUser(t *testing.T) {
	defer testPGStorage.clearTable(pgTableUser)

	// username does not exist
	func() {
		_, err := testPGStorage.AuthenticateUser("jaluhrman", "gabagool")
		if err == nil {
			t.Error("error should have been returned when username does not exist")
		}
	}()

	user := types.NewUser("jaluhrman", "password")
	testPGStorage.CreateUser(user)

	// bad password
	func() {
		_, err := testPGStorage.AuthenticateUser("jaluhrman", "wrong")
		if err == nil {
			t.Error("error should have been returned when password incorrect")
		}
	}()

	// success case
	func() {
		_, err := testPGStorage.AuthenticateUser("jaluhrman", "password")
		if err != nil {
			t.Error(err)
		}
	}()
}

func Test_CreateSession(t *testing.T) {
	defer testPGStorage.clearTable(pgTableSession)

	const (
		userID = 1
	)

	// success case
	func() {
		err := testPGStorage.CreateSession(types.NewSession(userID))
		if err != nil {
			t.Error(err)
		}
	}()

	// user id already exists
	func() {
		err := testPGStorage.CreateSession(types.NewSession(userID))
		if err == nil {
			t.Error("error should have been returned when session useID already exists")
		}
	}()
}

func Test_DeleteSessionByUserID(t *testing.T) {
	defer testPGStorage.clearTable(pgTableSession)

	const (
		userID = 5
	)

	// success case
	func() {
		testPGStorage.CreateSession(types.NewSession(userID))
		err := testPGStorage.DeleteSessionByUserID(userID)
		if err != nil {
			t.Error(err)
		}
	}()
}

func Test_DeleteSessionByToken(t *testing.T) {
	defer testPGStorage.clearTable(pgTableSession)

	const (
		userID = 3
	)

	// success case
	func() {
		s := types.NewSession(userID)
		token := s.Token

		testPGStorage.CreateSession(s)
		err := testPGStorage.DeleteSessionByToken(token)
		if err != nil {
			t.Error(err)
		}
	}()
}

func Test_AuthenticateSession(t *testing.T) {
	defer testPGStorage.clearTable(pgTableSession)

	// doesn't exist
	func() {
		_, err := testPGStorage.AuthenticateSession("random")
		if err == nil {
			t.Error("error should have been returned when session doesn't exist")
		}
	}()

	// success case
	func() {
		s := types.NewSession(1)
		testPGStorage.CreateSession(s)
		id, err := testPGStorage.AuthenticateSession(s.Token)
		if err != nil {
			t.Error(err)
		}
		if id != 1 {
			t.Error("returned user id was not correct")
		}
	}()
}

func Test_CreateExerciseType(t *testing.T) {
	defer testPGStorage.clearTable(pgTableExerciseType)

	testImage := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{200, 100},
	})

	exType := types.NewExerciseType("random name", testImage, types.Push, types.Quads)

	// success case
	func() {
		err := testPGStorage.CreateExerciseType(exType)
		if err != nil {
			t.Error(err)
		}
	}()

	// name already taken
	func() {
		err := testPGStorage.CreateExerciseType(exType)
		if err == nil {
			t.Error("Error should have been returned when exercise type name already taken")
		}
	}()
}

func Test_CreateWorkoutTemplate(t *testing.T) {
	// success case
	func() {
		wTemp := &types.WorkoutTemplate{
			UserID: 1,
			Name:   "test workout template",
		}

		var exTemps []types.ExerciseTemplate
		for i := 0; i < 3; i++ {
			exTemp := types.ExerciseTemplate{
				ExerciseTypeID: uint(i),
			}

			var sgTemps []types.SetGroupTemplate
			for j := 0; j < 3; j++ {
				sgTemp := types.SetGroupTemplate{
					Sets: uint(j),
					Reps: uint(j),
				}

				sgTemps = append(sgTemps, sgTemp)
			}

			exTemp.SetGroupTemplates = sgTemps
			exTemps = append(exTemps, exTemp)
		}

		wTemp.ExerciseTemplates = exTemps

		err := testPGStorage.CreateWorkoutTemplate(wTemp)
		if err != nil {
			t.Error(err)
		}

		for i, exTemp := range wTemp.ExerciseTemplates {
			if exTemp.WorkoutTemplateID != wTemp.ID {
				t.Errorf("Exercise template %d had workout template ID %d, should have been %d",
					i, exTemp.WorkoutTemplateID, wTemp.ID)
			}

			for j, sgTemp := range exTemp.SetGroupTemplates {
				if sgTemp.ExerciseTemplateID != exTemp.ID {
					t.Errorf("Setgroup template %d had exercise template ID %d, should have been %d",
						j, sgTemp.ExerciseTemplateID, exTemp.ID)
				}

				fmt.Println(sgTemp.ExerciseTemplateID, exTemp.ID)
			}
		}
	}()
}
