package main

import (
	"os"

	"github.com/joeluhrman/Lift-Tracker/db"
	"github.com/joeluhrman/Lift-Tracker/server"
)

func main() {
	var (
		isProd       bool
		dbConfig     *db.Config
		serverConfig *server.Config
	)

	if len(os.Args) > 1 {
		isProd = os.Args[1] == "-prod"
	}

	if isProd {
		dbConfig = db.TestDBConfig
		serverConfig = server.TestServerConfig
	} else {
		dbConfig = db.TestDBConfig
		serverConfig = server.TestServerConfig
	}

	db.MustConnect(dbConfig)
	defer db.MustClose()

	server.MustStart(serverConfig)
}
