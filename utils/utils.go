package utils

import (
	"os"
)

func MustReadFile(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}
