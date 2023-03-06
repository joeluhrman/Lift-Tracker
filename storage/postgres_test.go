package storage

import (
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

	tables = []string{pgTableUser, pgTableSession, pgTableExerciseType} /*pgTableLogSetgroup, pgTableLogExercise, pgTableLogWorkout*/
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
		user.HashedPassword, _ = HashPassword(user.Password)

		err := testPGStorage.CreateUser(user, false)
		if err != nil {
			t.Error(err)
		}
	}()

	// username already exists
	func() {
		user := types.NewUser("jaluhrman", "123")
		user.HashedPassword, _ = HashPassword(user.Password)

		err := testPGStorage.CreateUser(user, false)
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
	user.HashedPassword, _ = HashPassword(user.Password)
	testPGStorage.CreateUser(user, false)

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

	exType := types.NewExerciseType(1, true, "random name", testImage, types.Push, types.Quads)

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

/*
func Test_CreateLoggedWorkout(t *testing.T) {
	defer testPGStorage.clearAllTables()

	// success case
	func() {
		exercises := make([]*types.Exercise, 5)
		for i := 0; i < len(exercises); i++ {
			setgroups := make([]*types.Setgroup, 5)
			for i := 0; i < len(setgroups); i++ {
				setgroups[i] = types.NewSetgroup(0, i, i, i)
			}

			name := "exercise " + strconv.Itoa(i)
			notes := "notes " + strconv.Itoa(i)
			exercises[i] = types.NewExercise(0, name, setgroups, notes)
		}

		w := types.NewWorkout(1, "test_workout", time.Now(), exercises, "test notes")
		err := testPGStorage.CreateLoggedWorkout(w)
		if err != nil {
			t.Error(err)
		}
	}()
}
*/
