package testing

import (
	"os"
	"path"
	"runtime"
)

// Sets working directory of test suites to root so paths to resources
// are the same.
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
