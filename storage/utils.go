package storage

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

// things that I'm not sure which package they should go in

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func passwordMatchesHash(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func UsernameMeetsRequirements(username string) bool {
	return len(username) > 2
}

func PasswordMeetsRequirements(password string) bool {
	if password == "" {
		return false
	}

	return true
}

// not sure where to put this, can't put in main package because need for testing unfortunately
func MustReadFile(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}
