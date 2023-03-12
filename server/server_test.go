package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/go-cmp/cmp"
	"github.com/joeluhrman/Lift-Tracker/types"
)

const (
	wrongCodef = "code was %d, should have been %d"
)

var (
	// basic test server
	testServer = New("", &testStorage{}, middleware.Logger)

	// test server with middleware to set session for logged in tests
	testLoggedInServer = New("", &testStorage{},
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.AddCookie(types.NewSession(1).Cookie())
				ctx := context.WithValue(r.Context(), "user_id", 1)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		},
		middleware.Logger,
	)
)

type testStorage struct{}

func (t *testStorage) CreateUser(user *types.User) error {
	return nil
}

func (t *testStorage) AuthenticateUser(username string, password string) (int, error) {
	return 1, nil
}

func (t *testStorage) CreateSession(s *types.Session) error {
	return nil
}

func (t *testStorage) DeleteSessionByUserID(userID int) error {
	return nil
}

func (t *testStorage) DeleteSessionByToken(token string) error {
	return nil
}

func (t *testStorage) AuthenticateSession(token string) (int, error) {
	return 1, nil
}

func (t *testStorage) CreateExerciseType(exerciseType *types.ExerciseType) error {
	return nil
}

func (t *testStorage) GetExerciseTypes() ([]types.ExerciseType, error) {
	var eTypes []types.ExerciseType
	for i := 0; i < 5; i++ {
		eType := types.ExerciseType{
			ID:   uint(i + 1),
			Name: "eType " + strconv.Itoa(i+1),
			//Image:       image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{200, 100}}),
			PPLType:     types.Push,
			MuscleGroup: types.Calves,
		}

		eTypes = append(eTypes, eType)
	}

	return eTypes, nil
}

func (t *testStorage) CreateWorkoutTemplate(workoutTemplate *types.WorkoutTemplate) error {
	return nil
}

func (t *testStorage) GetWorkoutTemplates(userID uint) ([]types.WorkoutTemplate, error) {
	return nil, nil
}

func sendMockHTTPRequest(method string, endpoint string, data *bytes.Buffer, router http.Handler) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()

	if data == nil {
		data = &bytes.Buffer{}
	}

	req, err := http.NewRequest(method, endpoint, data)
	if err != nil {
		panic(err)
	}

	router.ServeHTTP(rec, req)

	return rec
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_CreateUserEndpoint(t *testing.T) {
	method := http.MethodPost
	endpoint := routeApiV1 + endUser
	successCode := http.StatusAccepted
	badJSONCode := http.StatusBadRequest
	badPasswordCode := http.StatusNotAcceptable

	// Bad JSON
	func() {
		rec := sendMockHTTPRequest(method, endpoint, nil, testServer.Handler)
		if rec.Code != badJSONCode {
			t.Errorf(wrongCodef, rec.Code, badJSONCode)
		}
	}()

	// Password doesn't meet requirements
	func() {
		user := types.NewUser("jaluhrman", "")

		json, _ := json.Marshal(user)
		body := bytes.NewBuffer(json)

		rec := sendMockHTTPRequest(method, endpoint, body, testServer.Handler)
		if rec.Code != badPasswordCode {
			t.Errorf(wrongCodef, rec.Code, badPasswordCode)
		}
	}()

	// Success case
	func() {
		user := types.NewUser("jaluhrman", "123")

		json, _ := json.Marshal(user)
		body := bytes.NewBuffer(json)

		rec := sendMockHTTPRequest(method, endpoint, body, testServer.Handler)
		if rec.Code != successCode {
			t.Errorf(wrongCodef, rec.Code, successCode)
		}
	}()
}

func Test_LoginEndpoint(t *testing.T) {
	method := http.MethodPost
	endpoint := routeApiV1 + endLogin
	badJSONCode := http.StatusBadRequest

	// bad json
	func() {
		rec := sendMockHTTPRequest(method, endpoint, nil, testServer.Handler)
		if rec.Code != badJSONCode {
			t.Errorf(wrongCodef, rec.Code, badJSONCode)
		}
	}()

	// success case
	func() {
		loginInfo := &types.User{
			Username:       "jaluhrman",
			HashedPassword: "123",
		}

		json, _ := json.Marshal(loginInfo)
		body := bytes.NewBuffer(json)

		rec := sendMockHTTPRequest(method, endpoint, body, testServer.Handler)

		// check correct response code
		if rec.Code != http.StatusOK {
			t.Errorf(wrongCodef, rec.Code, http.StatusOK)
		}

		// check session cookie has been set correctly
		sessionSet := false
		cookies := rec.Result().Cookies()
		for _, cookie := range cookies {
			if cookie.Name == types.SessionKey && cookie.Value != "" {
				sessionSet = true
			}
		}
		if !sessionSet {
			t.Error("session cookie was not set")
		}
	}()
}

func Test_LogoutEndpoint(t *testing.T) {
	const (
		method   = http.MethodPost
		endpoint = routeApiV1 + endLogout
		userID   = 1
	)

	// cookies reset correctly
	func() {
		rec := sendMockHTTPRequest(method, endpoint, nil, testLoggedInServer.Handler)
		if rec.Code != http.StatusOK {
			t.Errorf(wrongCodef, rec.Code, http.StatusOK)
		}

		cookieReset := false
		cookies := rec.Result().Cookies()
		for _, cookie := range cookies {
			if cookie.Name == types.SessionKey && cookie.Value == "" {
				cookieReset = true
			}
		}
		if !cookieReset {
			t.Error("session cookie was not reset")
		}
	}()
}

func Test_CreateWorkoutTemplateEndpoint(t *testing.T) {
	// success case
	func() {
		wTemp := &types.WorkoutTemplate{
			UserID: 1,
			Name:   "test wTemp",
			ExerciseTemplates: []types.ExerciseTemplate{
				{
					WorkoutTemplateID: 1,
					ExerciseTypeID:    1,
					SetGroupTemplates: []types.SetGroupTemplate{
						{
							ExerciseTemplateID: 1,
							Sets:               5,
							Reps:               5,
						},
					},
				},
			},
		}

		json, _ := json.Marshal(wTemp)
		data := bytes.NewBuffer(json)

		rec := sendMockHTTPRequest(http.MethodPost, routeApiV1+endWorkoutTemplate, data,
			testLoggedInServer.Handler)
		if rec.Code != http.StatusAccepted {
			t.Errorf(wrongCodef, rec.Code, http.StatusAccepted)
		}
	}()

	// user not logged in
	func() {
		rec := sendMockHTTPRequest(http.MethodPost, routeApiV1+endWorkoutTemplate, nil,
			testServer.Handler)
		if rec.Code != http.StatusUnauthorized {
			t.Errorf(wrongCodef, rec.Code, http.StatusUnauthorized)
		}
	}()

	// bad json
	func() {
		rec := sendMockHTTPRequest(http.MethodPost, routeApiV1+endWorkoutTemplate, nil,
			testLoggedInServer.Handler)
		if rec.Code != http.StatusBadRequest {
			t.Errorf(wrongCodef, rec.Code, http.StatusBadRequest)
		}
	}()
}

func Test_GetExerciseTypesEndpoint(t *testing.T) {
	// success case
	func() {
		eTypes, _ := testLoggedInServer.storage.GetExerciseTypes()

		rec := sendMockHTTPRequest(http.MethodGet, routeApiV1+endExerciseType, nil, testLoggedInServer.Handler)
		if rec.Code != http.StatusFound {
			t.Errorf(wrongCodef, rec.Code, http.StatusFound)
			return
		}

		var response []types.ExerciseType
		err := json.NewDecoder(rec.Body).Decode(&response)
		if err != nil {
			t.Error(err)
			return
		}

		if !cmp.Equal(eTypes, response) {
			t.Error(err)
		}
	}()

	// not logged in
	func() {
		rec := sendMockHTTPRequest(http.MethodGet, routeApiV1+endExerciseType, nil, testServer.Handler)
		if rec.Code != http.StatusUnauthorized {
			t.Errorf(wrongCodef, rec.Code, http.StatusUnauthorized)
		}
	}()
}
