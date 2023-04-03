package storage

import (
	"image"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/joeluhrman/Lift-Tracker/types"
)

var (
	testPGDriver = "pgx"
	testPGApiKey = string(MustReadFile("../api_key_test.txt"))
	testPGURL    = "postgresql://jaluhrman:" + testPGApiKey + "@db.bit.io/jaluhrman/Lift-Tracker-Test"

	testPGStorage = &testPostgres{
		Postgres: NewPostgres(testPGDriver, testPGURL),
	}

	tables = []string{pgTableSetGroupLog, pgTableSetGroupTemplate, pgTableExerciseLog,
		pgTableExerciseTemplate, pgTableExerciseType, pgTableWorkoutLog, pgTableWorkoutTemplate,
		pgTableSession, pgTableUser}
)

// wrapper for test methods to avoid confusion
type testPostgres struct {
	*Postgres
}

func (t *testPostgres) clearAllTables() {
	for _, table := range tables {
		t.clearTable(table)
	}
}

func (t *testPostgres) clearTable(tName string) {
	_, err := t.conn.Exec("DELETE FROM " + tName)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	testPGStorage.MustOpen()
	testPGStorage.clearAllTables()
	code := m.Run()
	testPGStorage.MustClose()
	os.Exit(code)
}

func Test_CreateUser(t *testing.T) {
	defer testPGStorage.clearTable(pgTableUser)

	// Success case
	func() {
		user := &types.User{Username: "jaluhrman", Email: "goober@goober.com"}
		password := "goober"

		err := testPGStorage.CreateUser(user, password)
		if err != nil {
			t.Error(err)
		}

		if user.ID == 0 {
			t.Error("Struct's ID was not updated")
		}
	}()

	// username already exists
	func() {
		user := &types.User{Username: "jaluhrman", Email: "goober@goober.com"}
		password := "123"

		err := testPGStorage.CreateUser(user, password)
		if err == nil {
			t.Error("should have produced an error when username already taken")
		}
	}()
}

func Test_GetUser(t *testing.T) {
	defer testPGStorage.clearTable(pgTableUser)

	// user doesn't exist
	func() {
		if _, err := testPGStorage.GetUser(1); err == nil {
			t.Error("error should have been returned when user doesn't exist")
		}
	}()

	// success case
	func() {
		user := &types.User{
			Username: "user",
			Email:    "user@user.com",
		}
		password := "password"

		testPGStorage.CreateUser(user, password)

		ret, err := testPGStorage.GetUser(user.ID)
		if err != nil {
			t.Error(err)
			return
		}

		user.HashedPassword = ""
		if !cmp.Equal(user, &ret) {
			t.Error("returned user was not the same as created")
			return
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

	user := &types.User{Username: "jaluhrman"}
	password := "password"
	testPGStorage.CreateUser(user, password)

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

		if exType.ID == 0 {
			t.Error("Struct's id was not updated")
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

func Test_GetExerciseTypes(t *testing.T) {
	// success case
	func() {
		testImage := image.NewRGBA(image.Rectangle{
			image.Point{0, 0},
			image.Point{200, 100},
		})

		exType := types.NewExerciseType("random name", testImage, types.Push, types.Quads)
		testPGStorage.CreateExerciseType(exType)

		rExTypes, err := testPGStorage.GetExerciseTypes()
		if err != nil {
			t.Error(err)
			return
		}

		if cmp.Equal(rExTypes[0], exType) {
			t.Error("Returned exercise type did not equal original")
		}
	}()
}

func Test_CreateWorkoutTemplate(t *testing.T) {
	defer testPGStorage.clearAllTables()

	// success case
	func() {
		user := &types.User{Username: "Bingus"}
		password := "Pringus"
		testPGStorage.CreateUser(user, password)

		testImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{200, 100}})
		exType := &types.ExerciseType{
			Name:        "test type",
			Image:       testImage,
			PPLType:     types.Push,
			MuscleGroup: types.Calves,
		}
		testPGStorage.CreateExerciseType(exType)

		wTemp := &types.WorkoutTemplate{
			UserID: uint(user.ID),
			Name:   "test workout template",
		}

		var exTemps []types.ExerciseTemplate
		for i := 0; i < 3; i++ {
			exTemp := types.ExerciseTemplate{
				ExerciseTypeID: exType.ID,
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

		if wTemp.ID <= 0 {
			t.Error("wTemp.ID was not scanned properly")
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
			}
		}
	}()
}

func Test_GetWorkoutTemplates(t *testing.T) {
	defer testPGStorage.clearAllTables()

	const loops = 3

	// success case
	func() {
		testUser := &types.User{Username: "jaluhrman"}
		password := "goober2000"
		testPGStorage.CreateUser(testUser, password)

		var wTemps []types.WorkoutTemplate
		for i := 0; i < loops; i++ {
			eType := types.NewExerciseType("type "+strconv.Itoa(i),
				image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{200, 100}}),
				types.Push, types.Abductors,
			)
			testPGStorage.CreateExerciseType(eType)

			wTemp := &types.WorkoutTemplate{
				UserID: uint(testUser.ID),
				Name:   "wTemp " + strconv.Itoa(i),
			}

			for j := 0; j < loops; j++ {
				eTemp := &types.ExerciseTemplate{
					WorkoutTemplateID: wTemp.ID,
					ExerciseTypeID:    eType.ID,
				}

				for k := 0; k < loops; k++ {
					sTemp := &types.SetGroupTemplate{
						ExerciseTemplateID: eTemp.ID,
						Sets:               uint(k),
						Reps:               uint(k),
					}

					eTemp.SetGroupTemplates = append(eTemp.SetGroupTemplates, *sTemp)
				}

				wTemp.ExerciseTemplates = append(wTemp.ExerciseTemplates, *eTemp)
			}

			testPGStorage.CreateWorkoutTemplate(wTemp)
			wTemps = append(wTemps, *wTemp)
		}

		rTemps, err := testPGStorage.GetWorkoutTemplates(uint(testUser.ID))
		if err != nil {
			t.Error(err)
		}

		if len(rTemps) != len(wTemps) {
			t.Errorf("returned slice had length %d, original has length %d", len(rTemps), len(wTemps))
		}

		if !(cmp.Equal(rTemps, wTemps)) {
			t.Error("original and returned slices were not equal")
		}
	}()
}

func Test_CreateWorkoutLog(t *testing.T) {
	defer testPGStorage.clearAllTables()

	// success case
	func() {
		user := &types.User{Username: "spoingus"}
		password := "moingus"
		testPGStorage.CreateUser(user, password)

		testImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{200, 100}})
		exType := &types.ExerciseType{
			Name:        "test type",
			Image:       testImage,
			PPLType:     types.Push,
			MuscleGroup: types.Calves,
		}
		testPGStorage.CreateExerciseType(exType)

		wLog := &types.WorkoutLog{
			UserID: user.ID,
			Date:   time.Time{},
			Name:   "test workout template",
			Notes:  "test notes",
		}

		var eLogs []types.ExerciseLog
		for i := 0; i < 3; i++ {
			eLog := types.ExerciseLog{
				ExerciseTypeID: exType.ID,
				Notes:          "test notes",
			}

			var sgLogs []types.SetGroupLog
			for j := 0; j < 3; j++ {
				sgLog := types.SetGroupLog{
					Sets:   uint(j),
					Reps:   uint(j),
					Weight: float32(i * j),
				}

				sgLogs = append(sgLogs, sgLog)
			}

			eLog.SetGroupLogs = sgLogs
			eLogs = append(eLogs, eLog)
		}

		wLog.ExerciseLogs = eLogs

		err := testPGStorage.CreateWorkoutLog(wLog)
		if err != nil {
			t.Error(err)
			return
		}

		if wLog.ID <= 0 {
			t.Error("wLog.ID was not scanned correctly")
			return
		}

		for i, eLog := range wLog.ExerciseLogs {
			if eLog.WorkoutLogID != wLog.ID {
				t.Errorf("Exercise template %d had workout template ID %d, should have been %d",
					i, eLog.WorkoutLogID, wLog.ID)
			}

			for j, sgLog := range eLog.SetGroupLogs {
				if sgLog.ExerciseLogID != eLog.ID {
					t.Errorf("Setgroup template %d had exercise template ID %d, should have been %d",
						j, sgLog.ExerciseLogID, eLog.ID)
				}
			}
		}
	}()
}

func Test_GetWorkoutLogs(t *testing.T) {
	// success case
	func() {
		const loops = 3

		testUser := &types.User{Username: "jaluhrman"}
		password := "goober2000"
		testPGStorage.CreateUser(testUser, password)

		var wLogs []types.WorkoutLog
		for i := 0; i < loops; i++ {
			eType := types.NewExerciseType("type "+strconv.Itoa(i),
				image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{200, 100}}),
				types.Push, types.Abductors,
			)
			testPGStorage.CreateExerciseType(eType)

			wLog := &types.WorkoutLog{
				UserID: uint(testUser.ID),
				Date:   time.Time{},
				Name:   "wTemp " + strconv.Itoa(i),
				Notes:  "test",
			}

			for j := 0; j < loops; j++ {
				eLog := types.ExerciseLog{
					WorkoutLogID:   wLog.ID,
					ExerciseTypeID: eType.ID,
					Notes:          "test notes",
				}

				for k := 0; k < loops; k++ {
					sLog := types.SetGroupLog{
						ExerciseLogID: eLog.ID,
						Sets:          uint(k),
						Reps:          uint(k),
						Weight:        float32(k),
					}

					eLog.SetGroupLogs = append(eLog.SetGroupLogs, sLog)
				}

				wLog.ExerciseLogs = append(wLog.ExerciseLogs, eLog)
			}

			testPGStorage.CreateWorkoutLog(wLog)
			wLogs = append(wLogs, *wLog)
		}

		rLogs, err := testPGStorage.GetWorkoutLogs(testUser.ID)
		if err != nil {
			t.Error(err)
			return
		}

		if !cmp.Equal(wLogs, rLogs) {
			t.Error("returned logs were not equal to original")
			return
		}
	}()
}
