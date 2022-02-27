package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// TESTING FUNCTION ONLY
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// just testing stuff for now

	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":8080", nil))

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
