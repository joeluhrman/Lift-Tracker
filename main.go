package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// just testing stuff for now

	// testing static html files
	port := os.Getenv("PORT")
	fs := http.FileServer(http.Dir("./web/templates"))
	http.Handle("/", fs)

	log.Println("Listening on :" + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
