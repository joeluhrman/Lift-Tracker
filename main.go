package main

import (
	"fmt"
	"log"
	"net/http"
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
	fs := http.FileServer(http.Dir("./web/templates"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
